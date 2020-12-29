package static_localized

import (
	"fmt"
	"net"
	"net/url"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	envoyauth "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	"github.com/envoyproxy/go-control-plane/pkg/wellknown"
	pbgostruct "github.com/golang/protobuf/ptypes/struct"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1static_localized "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static_localized"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/solo-kit/pkg/errors"
)

const (
	// TODO: make solo-projects use this constant
	TransportSocketMatchKey = "envoy.transport_socket_match"

	HttpPathCheckerName = "io.solo.health_checkers.http_path"
	PathFieldName       = "path"
)

var _ plugins.Plugin = new(plugin)
var _ plugins.UpstreamPlugin = new(plugin)

type plugin struct{}

func NewPlugin() plugins.Plugin {
	return &plugin{}
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	staticSpec, ok := u.UpstreamType.(*v1.Upstream_StaticLocalized)
	if !ok {
		return nil, nil
	}
	if len(staticSpec.StaticLocalized.LocalizedHosts) == 0 {
		return nil, errors.Errorf("must provide at least 1 localized_host in static localized spec")
	}

	if len(staticSpec.StaticLocalized.LocalizedHosts[0].Hosts) == 0 {
		return nil, errors.Errorf("must provide at least 1 host in static localized spec")
	}

	hostToResolve := staticSpec.StaticLocalized.LocalizedHosts[0].Hosts[0]
	return url.Parse(fmt.Sprintf("tcp://%v:%v", hostToResolve.Addr, hostToResolve.Port))
}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoy_config_cluster_v3.Cluster) error {
	staticSpec, ok := in.UpstreamType.(*v1.Upstream_StaticLocalized)
	if !ok {
		// not ours
		return nil
	}

	spec := staticSpec.StaticLocalized
	var foundSslPort bool
	var hostname string

	out.ClusterDiscoveryType = &envoy_config_cluster_v3.Cluster_Type{
		Type: envoy_config_cluster_v3.Cluster_STATIC,
	}

	for localityIdx, localizedHosts := range spec.LocalizedHosts {
		for _, host := range localizedHosts.Hosts {
			if host.Addr == "" {
				return errors.Errorf("addr cannot be empty for host")
			}
			if host.Port == 0 {
				return errors.Errorf("port cannot be empty for host")
			}
			if host.Port == 443 {
				foundSslPort = true
			}
			ip := net.ParseIP(host.Addr)
			if ip == nil {
				// can't parse ip so this is a dns hostname.
				// save the first hostname for use with sni
				if hostname == "" {
					hostname = host.Addr
				}
			}

			if out.LoadAssignment == nil {
				out.LoadAssignment = &envoy_config_endpoint_v3.ClusterLoadAssignment{
					ClusterName: out.Name,
					Endpoints:   []*envoy_config_endpoint_v3.LocalityLbEndpoints{{}},
				}
			}

			// set the locality, if not already set
			if out.LoadAssignment.Endpoints[localityIdx].Locality == nil {
				out.LoadAssignment.Endpoints[localityIdx].Locality = getLocalityForHost(localizedHosts.Locality)
			}

			// append the lb endpoint
			out.LoadAssignment.Endpoints[localityIdx].LbEndpoints = append(out.LoadAssignment.Endpoints[localityIdx].LbEndpoints,
				getLbEndpointForHost(spec, host))

			// the priority is not set, and therefore treated as P=0
		}
	}

	// if host port is 443 or if the user wants it, we will use TLS
	if spec.UseTls || foundSslPort {
		// tell envoy to use TLS to connect to this upstream
		// TODO: support client certificates
		if out.TransportSocket == nil {
			tlsContext := &envoyauth.UpstreamTlsContext{
				// TODO(yuval-k): Add verification context
				Sni: hostname,
			}
			out.TransportSocket = &envoy_config_core_v3.TransportSocket{
				Name:       wellknown.TransportSocketTls,
				ConfigType: &envoy_config_core_v3.TransportSocket_TypedConfig{TypedConfig: utils.MustMessageToAny(tlsContext)},
			}
		}
	}
	if out.TransportSocket != nil {
		for _, localizedHosts := range spec.LocalizedHosts {
			for _, host := range localizedHosts.Hosts {
				sniname := sniAddr(spec, host)
				if sniname == "" {
					continue
				}
				ts, err := mutateSni(out.TransportSocket, sniname)
				if err != nil {
					return err
				}
				out.TransportSocketMatches = append(out.TransportSocketMatches, &envoy_config_cluster_v3.Cluster_TransportSocketMatch{
					Name:            name(spec, host),
					Match:           metadataMatch(spec, host),
					TransportSocket: ts,
				})
			}
		}
	}

	// the upstream has a DNS name. We need to resolve the DNS name
	if hostname != "" {
		// TODO (sam-heilbron)
	}

	return nil
}

func getLbEndpointForHost(spec *v1static_localized.UpstreamSpec, host *v1static_localized.Host) *envoy_config_endpoint_v3.LbEndpoint {
	return &envoy_config_endpoint_v3.LbEndpoint{
		Metadata: getMetadata(spec, host),
		HostIdentifier: &envoy_config_endpoint_v3.LbEndpoint_Endpoint{
			Endpoint: &envoy_config_endpoint_v3.Endpoint{
				Hostname: host.Addr,
				Address: &envoy_config_core_v3.Address{
					Address: &envoy_config_core_v3.Address_SocketAddress{
						SocketAddress: &envoy_config_core_v3.SocketAddress{
							Protocol: envoy_config_core_v3.SocketAddress_TCP,
							Address:  host.Addr,
							PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
								PortValue: host.Port,
							},
						},
					},
				},
				HealthCheckConfig: &envoy_config_endpoint_v3.Endpoint_HealthCheckConfig{
					Hostname: host.Addr,
				},
			},
		},
	}
}

func getLocalityForHost(locality *v1static_localized.Locality) *envoy_config_core_v3.Locality {
	if locality == nil {
		return nil
	}
	return &envoy_config_core_v3.Locality{
		Region:  locality.GetRegion(),
		Zone:    locality.GetZone(),
		SubZone: locality.GetSubZone(),
	}
}

func mutateSni(in *envoy_config_core_v3.TransportSocket, sni string) (*envoy_config_core_v3.TransportSocket, error) {
	copy := *in

	// copy the sni
	cfg, err := utils.AnyToMessage(copy.GetTypedConfig())
	if err != nil {
		return nil, err
	}

	typedCfg, ok := cfg.(*envoyauth.UpstreamTlsContext)
	if !ok {
		return nil, errors.Errorf("unknown tls config type: %T", cfg)
	}
	typedCfg.Sni = sni

	copy.ConfigType = &envoy_config_core_v3.TransportSocket_TypedConfig{TypedConfig: utils.MustMessageToAny(typedCfg)}

	return &copy, nil
}

func sniAddr(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) string {
	if in.GetSniAddr() != "" {
		return in.GetSniAddr()
	}
	if spec.GetAutoSniRewrite() == nil || spec.GetAutoSniRewrite().GetValue() {
		return in.GetAddr()
	}
	return ""
}

func getMetadata(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) *envoy_config_core_v3.Metadata {
	if in == nil {
		return nil
	}
	var meta *envoy_config_core_v3.Metadata
	sniaddr := sniAddr(spec, in)
	if sniaddr != "" {
		if meta == nil {
			meta = &envoy_config_core_v3.Metadata{FilterMetadata: map[string]*pbgostruct.Struct{}}
		}
		meta.FilterMetadata[TransportSocketMatchKey] = metadataMatch(spec, in)
	}

	if in.GetHealthCheckConfig().GetPath() != "" {
		if meta == nil {
			meta = &envoy_config_core_v3.Metadata{FilterMetadata: map[string]*pbgostruct.Struct{}}
		}
		meta.FilterMetadata[HttpPathCheckerName] = &pbgostruct.Struct{
			Fields: map[string]*pbgostruct.Value{
				PathFieldName: {
					Kind: &pbgostruct.Value_StringValue{
						StringValue: in.GetHealthCheckConfig().GetPath(),
					},
				},
			},
		}

	}
	return meta
}

func name(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) string {
	return fmt.Sprintf("%s;%s:%d", sniAddr(spec, in), in.Addr, in.Port)
}

func metadataMatch(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) *pbgostruct.Struct {
	return &pbgostruct.Struct{
		Fields: map[string]*pbgostruct.Value{
			name(spec, in): {
				Kind: &pbgostruct.Value_BoolValue{
					BoolValue: true,
				},
			},
		},
	}
}
