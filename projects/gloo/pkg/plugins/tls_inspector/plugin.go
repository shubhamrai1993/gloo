package tls_inspector

import (
	envoy_config_listener_v3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"

	envoy_tls_inspector "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"
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
	if in.GetSslConfigurations() != nil {
		configEnvoy := &envoy_tls_inspector.TlsInspector{}
		msg, err := utils.MessageToAny(configEnvoy)
		if err == nil {
			return nil
		}
		TLSInspector := &envoy_config_listener_v3.ListenerFilter{
			Name: wellknown.TlsInspector,
			ConfigType: &envoy_config_listener_v3.ListenerFilter_TypedConfig{
				TypedConfig: msg,
			},
		}

		out.ListenerFilters = append(out.ListenerFilters, TLSInspector)
	}

	return nil
}
