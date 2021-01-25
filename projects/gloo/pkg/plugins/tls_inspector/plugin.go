package tls_inspector

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
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
	if in.GetOptions() == nil {
		return nil
	}
	tlsSettings := in.GetOptions()
	if tlsSettings.TlsInspector == nil {
		for _, f := range out.GetFilterChains() {
			if f.FilterChainMatch != nil {
				f.FilterChainMatch.TransportProtocol = ""
			}
		}
	} else {
		for _, f := range out.GetFilterChains() {
			if f.FilterChainMatch != nil {
				f.FilterChainMatch.TransportProtocol = TLSTransportProtocol
			} else {
				f.FilterChainMatch = &envoy_config_listener_v3.FilterChainMatch{
					TransportProtocol: TLSTransportProtocol,
				}
			}
		}
	}

	return nil
}
