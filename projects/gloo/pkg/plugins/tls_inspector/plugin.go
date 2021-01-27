package tls_inspector

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	envoy_tls_inspector "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
)

// filter should be called after routing decision has been made
var (
	pluginStage          = plugins.DuringStage(plugins.RouteStage)
	TLSTransportProtocol = "tls"
)

func NewPlugin() *plugin {
	return &plugin{}
}

var (
	_ plugins.Plugin         = new(plugin)
	_ plugins.ListenerPlugin = new(plugin)
)

type plugin struct{}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) ProcessListener(params plugins.Params, in *v1.Listener, out *envoy_config_listener_v3.Listener) error {
	configEnvoy := &envoy_tls_inspector.TlsInspector{}
	msg, err := utils.MessageToAny(configEnvoy)
	if err != nil {
		return err
	}
	tlsInspector := &envoy_config_listener_v3.ListenerFilter{
		Name: wellknown.TlsInspector,
		ConfigType: &envoy_config_listener_v3.ListenerFilter_TypedConfig{
			TypedConfig: msg,
		},
	}

	switch in.GetListenerType().(type) {
	case *v1.Listener_HttpListener:
		// automatically add tls inspector when ssl is enabled
		if in.GetSslConfigurations() != nil {
			out.ListenerFilters = append([]*envoy_config_listener_v3.ListenerFilter{tlsInspector}, out.ListenerFilters...)
		}
	case *v1.Listener_TcpListener:
		tcpListener := in.GetTcpListener()
		var hostHasConfig, sniCluster, sniMatch bool
		for _, host := range tcpListener.GetTcpHosts() {
			if host.GetSslConfig() != nil {
				hostHasConfig = true
			}
			if len(host.GetSslConfig().GetSniDomains()) > 0 {
				sniMatch = true
			}
			if host.GetDestination().GetForwardSniClusterName() != nil {
				sniCluster = true
			}
		}

		// If there is a forward SNI cluster, and no SNI matches, prepend the TLS inspector manually.
		if sniCluster && !sniMatch {
			out.ListenerFilters = append(
				[]*envoy_config_listener_v3.ListenerFilter{{Name: wellknown.TlsInspector}},
				out.ListenerFilters...,
			)
		} else if hostHasConfig || in.GetSslConfigurations() != nil {
			// for other cases, add TLS inspector
			out.ListenerFilters = append([]*envoy_config_listener_v3.ListenerFilter{tlsInspector}, out.ListenerFilters...)
		}

	}

	return nil
}
