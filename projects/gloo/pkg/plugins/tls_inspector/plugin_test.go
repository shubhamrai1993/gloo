package tls_inspector

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoyhttp "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tls_inspector"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	translatorutil "github.com/solo-io/gloo/projects/gloo/pkg/translator"
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

			filters := []*envoy_config_listener_v3.Filter{{
				Name: wellknown.HTTPConnectionManager,
			}}

			outl := &envoy_config_listener_v3.Listener{
				FilterChains: []*envoy_config_listener_v3.FilterChain{{
					Filters: filters,
				}},
			}

			p := NewPlugin()
			err := p.ProcessListener(params, in, outl)
			Expect(err).NotTo(HaveOccurred())

			var cfg envoyhttp.HttpConnectionManager
			err = translatorutil.ParseTypedConfig(filters[0], &cfg)
			Expect(err).NotTo(HaveOccurred())


		})

		It("tls inspector is ignored", func() {
			hl := &v1.HttpListener{}

			in := &v1.Listener{
				ListenerType: &v1.Listener_HttpListener{
					HttpListener: hl,
				},
				Options: &v1.ListenerOptions{},
			}

			filterChainMatch := &envoy_config_listener_v3.FilterChainMatch{}

			outl := &envoy_config_listener_v3.Listener{
				FilterChains: []*envoy_config_listener_v3.FilterChain{{
					FilterChainMatch: filterChainMatch,
				}},
			}

			p := NewPlugin()
			err := p.ProcessListener(params, in, outl)
			Expect(err).NotTo(HaveOccurred())

			for _, f := range outl.GetFilterChains() {
				Expect(f.FilterChainMatch.TransportProtocol).To(Equal(""))
			}
		})
	})
})
