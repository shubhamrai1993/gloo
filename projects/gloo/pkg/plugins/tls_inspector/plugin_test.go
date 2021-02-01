package tls_inspector

import (
	. "github.com/onsi/ginkgo"
)

var _ = Describe("Plugin", func() {
	//Context("tls inspector for http", func() {
	//
	//	var (
	//		params plugins.Params
	//	)
	//
	//	BeforeEach(func() {
	//		params = plugins.Params{}
	//	})
	//
	//	It("tls inspector is added", func() {
	//		hl := &v1.HttpListener{}
	//		in := &v1.Listener{
	//			ListenerType: &v1.Listener_HttpListener{
	//				HttpListener: hl,
	//			},
	//			SslConfigurations: []*v1.SslConfig{},
	//		}
	//
	//		filters := []*envoy_config_listener_v3.Filter{{}}
	//
	//		out := &envoy_config_listener_v3.Listener{
	//			FilterChains: []*envoy_config_listener_v3.FilterChain{{
	//				Filters: filters,
	//			}},
	//		}
	//
	//		p := NewPlugin()
	//		err := p.ProcessListener(params, in, out)
	//		Expect(err).NotTo(HaveOccurred())
	//
	//		configEnvoy := &envoy_tls_inspector.TlsInspector{}
	//		config, err := utils.MessageToAny(configEnvoy)
	//
	//		Expect(out.ListenerFilters).To(HaveLen(1))
	//		Expect(out.ListenerFilters[0].GetName()).To(Equal(wellknown.TlsInspector))
	//		Expect(out.ListenerFilters[0].GetTypedConfig()).To(Equal(config))
	//
	//	})
	//
	//	It("tls inspector is ignored", func() {
	//
	//		hl := &v1.HttpListener{}
	//		in := &v1.Listener{
	//			ListenerType: &v1.Listener_HttpListener{
	//				HttpListener: hl,
	//			},
	//		}
	//
	//		filters := []*envoy_config_listener_v3.Filter{{}}
	//
	//		out := &envoy_config_listener_v3.Listener{
	//			FilterChains: []*envoy_config_listener_v3.FilterChain{{
	//				Filters: filters,
	//			}},
	//		}
	//
	//		p := NewPlugin()
	//		err := p.ProcessListener(params, in, out)
	//		Expect(err).NotTo(HaveOccurred())
	//
	//		Expect(out.ListenerFilters).To(HaveLen(0))
	//	})
	//})
	//
	//Context("tls inspector for tcp", func() {
	//
	//	var (
	//		params plugins.Params
	//	)
	//
	//	BeforeEach(func() {
	//		params = plugins.Params{}
	//	})
	//
	//	It("SslConfigurations is set, tls inspector is added", func() {
	//		tl := &v1.TcpListener{}
	//		in := &v1.Listener{
	//			ListenerType: &v1.Listener_TcpListener{
	//				TcpListener: tl,
	//			},
	//			SslConfigurations: []*v1.SslConfig{},
	//		}
	//
	//		filters := []*envoy_config_listener_v3.Filter{{}}
	//
	//		out := &envoy_config_listener_v3.Listener{
	//			FilterChains: []*envoy_config_listener_v3.FilterChain{{
	//				Filters: filters,
	//			}},
	//		}
	//
	//		p := NewPlugin()
	//		err := p.ProcessListener(params, in, out)
	//		Expect(err).NotTo(HaveOccurred())
	//
	//		configEnvoy := &envoy_tls_inspector.TlsInspector{}
	//		config, err := utils.MessageToAny(configEnvoy)
	//
	//		Expect(out.ListenerFilters).To(HaveLen(1))
	//		Expect(out.ListenerFilters[0].GetName()).To(Equal(wellknown.TlsInspector))
	//		Expect(out.ListenerFilters[0].GetTypedConfig()).To(Equal(config))
	//	})
	//
	//	It("Tcp Host Ssl Config is set, tls inspector is added", func() {
	//		thost := &v1.TcpHost{}
	//		tl := &v1.TcpListener{
	//			TcpHosts: []*v1.TcpHost{thost},
	//		}
	//		in := &v1.Listener{
	//			ListenerType: &v1.Listener_TcpListener{
	//				TcpListener: tl,
	//			},
	//			SslConfigurations: []*v1.SslConfig{},
	//		}
	//
	//		filters := []*envoy_config_listener_v3.Filter{{}}
	//
	//		out := &envoy_config_listener_v3.Listener{
	//			FilterChains: []*envoy_config_listener_v3.FilterChain{{
	//				Filters: filters,
	//			}},
	//		}
	//
	//		p := NewPlugin()
	//		err := p.ProcessListener(params, in, out)
	//		Expect(err).NotTo(HaveOccurred())
	//
	//		configEnvoy := &envoy_tls_inspector.TlsInspector{}
	//		config, err := utils.MessageToAny(configEnvoy)
	//
	//		Expect(out.ListenerFilters).To(HaveLen(1))
	//		Expect(out.ListenerFilters[0].GetName()).To(Equal(wellknown.TlsInspector))
	//		Expect(out.ListenerFilters[0].GetTypedConfig()).To(Equal(config))
	//	})
	//
	//	It("tls inspector is ignored", func() {
	//		tl := &v1.TcpListener{}
	//		in := &v1.Listener{
	//			ListenerType: &v1.Listener_TcpListener{
	//				TcpListener: tl,
	//			},
	//		}
	//
	//		filters := []*envoy_config_listener_v3.Filter{{}}
	//
	//		out := &envoy_config_listener_v3.Listener{
	//			FilterChains: []*envoy_config_listener_v3.FilterChain{{
	//				Filters: filters,
	//			}},
	//		}
	//
	//		p := NewPlugin()
	//		err := p.ProcessListener(params, in, out)
	//		Expect(err).NotTo(HaveOccurred())
	//
	//		Expect(out.ListenerFilters).To(HaveLen(0))
	//	})
	//
	//	It("will prepend the TlsInspector when NO ServerName match present", func() {
	//		snap := &v1.ApiSnapshot{}
	//		out := &envoy_config_listener_v3.Listener{}
	//		tcpListener := &v1.TcpListener{
	//			TcpHosts: []*v1.TcpHost{
	//				{
	//					Name: "one",
	//					Destination: &v1.TcpHost_TcpAction{
	//						Destination: &v1.TcpHost_TcpAction_ForwardSniClusterName{
	//							ForwardSniClusterName: &empty.Empty{},
	//						},
	//					},
	//				},
	//			},
	//		}
	//		listener := &v1.Listener{
	//			ListenerType: &v1.Listener_TcpListener{
	//				TcpListener: tcpListener,
	//			},
	//		}
	//
	//		configEnvoy := &envoy_tls_inspector.TlsInspector{}
	//		config, err := utils.MessageToAny(configEnvoy)
	//
	//		p := NewPlugin()
	//		err = p.ProcessListener(plugins.Params{Snapshot: snap}, listener, out)
	//		Expect(err).NotTo(HaveOccurred())
	//		Expect(out.ListenerFilters).To(HaveLen(1))
	//		Expect(out.ListenerFilters[0].GetName()).To(Equal(wellknown.TlsInspector))
	//		Expect(out.ListenerFilters[0].GetTypedConfig()).To(Equal(config))
	//	})
	//})
})
