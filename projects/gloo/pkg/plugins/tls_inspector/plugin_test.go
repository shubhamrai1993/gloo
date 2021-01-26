package tls_inspector

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_tls_inspector "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tls_inspector"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
)

var _ = Describe("Plugin", func() {
	var (
		tlsConfig *tls_inspector.TlsInspector
	)
	Context("tls inspector", func() {

		var (
			params plugins.Params
		)

		BeforeEach(func() {
			tlsConfig = &tls_inspector.TlsInspector{}
			params = plugins.Params{}
		})

		It("tls inspector is added", func() {
			hl := &v1.HttpListener{}

			in := &v1.Listener{
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: hl,
				},
				Options: &v1.ListenerOptions{
					TlsInspector: tlsConfig,
				},
			}

			filters := []*envoy_config_listener_v3.Filter{{}}

			outl := &envoy_config_listener_v3.Listener{
				FilterChains: []*envoy_config_listener_v3.FilterChain{{
					Filters: filters,
				}},
			}

			p := NewPlugin()
			err := p.ProcessListener(params, in, outl)
			Expect(err).NotTo(HaveOccurred())

			configEnvoy := &envoy_tls_inspector.TlsInspector{}
			config, err := utils.MessageToAny(configEnvoy)

			for _, f := range outl.GetListenerFilters() {
				if f.Name == wellknown.TlsInspector {
					Expect(f.ConfigType).To(Equal(config))
				}
			}

		})

		It("tls inspector is ignored", func() {
			hl := &v1.HttpListener{}

			in := &v1.Listener{
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: hl,
				},
				Options: &v1.ListenerOptions{},
			}

			filters := []*envoy_config_listener_v3.Filter{{}}

			outl := &envoy_config_listener_v3.Listener{
				FilterChains: []*envoy_config_listener_v3.FilterChain{{
					Filters: filters,
				}},
			}

			p := NewPlugin()
			err := p.ProcessListener(params, in, outl)
			Expect(err).NotTo(HaveOccurred())

			configEnvoy := &envoy_tls_inspector.TlsInspector{}
			config, err := utils.MessageToAny(configEnvoy)

			for _, f := range outl.GetListenerFilters() {
				if f.Name == wellknown.TlsInspector {
					Expect(f.ConfigType).To(Equal(config))
				}
			}
		})
	})
})
