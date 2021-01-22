package envoy_test

import (
	"context"
	"os"
	"testing"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"k8s.io/client-go/rest"

	"github.com/solo-io/go-utils/log"
)

func TestE2e(t *testing.T) {
	if os.Getenv("KUBE2E_TESTS") != "envoy" {
		log.Warnf("This test is disabled. " +
			"To enable, set KUBE2E_TESTS to 'envoy' in your env.")
		return
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}

var (
	ctx    context.Context
	cancel context.CancelFunc
	err    error
	cfg    *rest.Config

	upstreamClient       gloov1.UpstreamClient
	virtualServiceClient gatewayv1.VirtualServiceClient
)

var _ = SynchronizedBeforeSuite(func() []byte {
	ctx, cancel = context.WithCancel(context.Background())
	cfg, err = kubeutils.GetConfig("", "")
	Expect(err).NotTo(HaveOccurred())

	cache := kube.NewKubeCache(ctx)
	upstreamClientFactory := &factory.KubeResourceClientFactory{
		Crd:         gloov1.UpstreamCrd,
		Cfg:         cfg,
		SharedCache: cache,
	}
	virtualServiceClientFactory := &factory.KubeResourceClientFactory{
		Crd:         gatewayv1.VirtualServiceCrd,
		Cfg:         cfg,
		SharedCache: cache,
	}
	upstreamClient, err = gloov1.NewUpstreamClient(ctx, upstreamClientFactory)
	Expect(err).NotTo(HaveOccurred())
	virtualServiceClient, err = gatewayv1.NewVirtualServiceClient(ctx, virtualServiceClientFactory)
	Expect(err).NotTo(HaveOccurred())

	// Wait for Virtual Service to be accepted
	Eventually(func() bool {
		vs, err := virtualServiceClient.Read(defaults.GlooSystem, "default", clients.ReadOpts{})
		Expect(err).NotTo(HaveOccurred())
		return vs.Status.GetState() == core.Status_Accepted
	}, "15s", "0.5s").Should(BeTrue())

	return nil
}, func([]byte) {})

var _ = SynchronizedAfterSuite(func() {}, func() {
	cancel()
})
