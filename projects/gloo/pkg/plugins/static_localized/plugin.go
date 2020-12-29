package static_localized

import (
	"context"
	"fmt"
	"net"
	"net/url"

	"github.com/rotisserie/eris"

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

type plugin struct {
	resolver DnsResolver
}

func NewPlugin(resolver DnsResolver) plugins.Plugin {
	return &plugin{
		resolver: resolver,
	}
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	staticSpec, ok := u.UpstreamType.(*v1.Upstream_StaticLocalized)
	if !ok {
		return nil, nil
	}
	spec := staticSpec.StaticLocalized

	if len(spec.LocalizedHosts) == 0 {
		return nil, errors.Errorf("must provide at least 1 localized_host in static localized spec")
	}

	if len(spec.LocalizedHosts[0].Hosts) == 0 {
		return nil, errors.Errorf("must provide at least 1 host in static localized spec")
	}

	hostToResolve := spec.LocalizedHosts[0].Hosts[0]
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
	var sniHostname string

	out.ClusterDiscoveryType = &envoy_config_cluster_v3.Cluster_Type{
		Type: envoy_config_cluster_v3.Cluster_STATIC,
	}

	for localityIdx, localizedHosts := range spec.LocalizedHosts {
		var lbEndpoints []*envoy_config_endpoint_v3.LbEndpoint
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
			if ip != nil {
				// successfully parsed ip, create an lb endpoint for this ip address
				lbEndpoints = append(lbEndpoints, getLbEndpointForIPHost(spec, host))
			} else {
				// can't parse ip so this is a dns hostname.
				// save the first hostname for use with sni
				if sniHostname == "" {
					sniHostname = host.Addr
				}

				// resolve the ip addresses of this dns hostname, create an lb endpoint for each ip address
				endpoints, err := getLbEndpointsForDnsHost(spec, host, p.resolver)
				if err != nil {
					return err
				}
				lbEndpoints = append(lbEndpoints, endpoints...)
			}
		}

		if len(lbEndpoints) == 0 {
			// avoid writing empty config
			continue
		}

		if out.LoadAssignment == nil {
			out.LoadAssignment = &envoy_config_endpoint_v3.ClusterLoadAssignment{
				ClusterName: out.Name,
				Endpoints:   []*envoy_config_endpoint_v3.LocalityLbEndpoints{{}},
			}
		}

		out.LoadAssignment.Endpoints[localityIdx] = &envoy_config_endpoint_v3.LocalityLbEndpoints{
			Priority:    0,
			Locality:    getLocalityForHost(localizedHosts.Locality),
			LbEndpoints: lbEndpoints,
		}
	}

	// if host port is 443 or if the user wants it, we will use TLS
	if spec.UseTls || foundSslPort {
		// tell envoy to use TLS to connect to this upstream
		if out.TransportSocket == nil {
			out.TransportSocket = getTransportSocket(sniHostname)
		}
	}

	if out.TransportSocket != nil {
		transportSocketMatches, err := getTransportSocketMatches(spec, out.TransportSocket)
		if err != nil {
			return err
		}
		out.TransportSocketMatches = transportSocketMatches
	}

	return nil
}

func getLbEndpointForIPHost(spec *v1static_localized.UpstreamSpec, host *v1static_localized.Host) *envoy_config_endpoint_v3.LbEndpoint {
	return &envoy_config_endpoint_v3.LbEndpoint{
		Metadata: getMetadataForHost(spec, host),
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

func getLbEndpointsForDnsHost(spec *v1static_localized.UpstreamSpec, host *v1static_localized.Host, resolver DnsResolver) ([]*envoy_config_endpoint_v3.LbEndpoint, error) {
	ipAddresses, err := resolveIpAddresses(context.TODO(), host.Addr, resolver)
	if err != nil {
		return nil, err
	}

	var lbEndpoints []*envoy_config_endpoint_v3.LbEndpoint

	for _, ipAddress := range ipAddresses {
		lbEndpoint := &envoy_config_endpoint_v3.LbEndpoint{
			Metadata: getMetadataForHost(spec, host),
			HostIdentifier: &envoy_config_endpoint_v3.LbEndpoint_Endpoint{
				Endpoint: &envoy_config_endpoint_v3.Endpoint{
					Hostname: ipAddress,
					Address: &envoy_config_core_v3.Address{
						Address: &envoy_config_core_v3.Address_SocketAddress{
							SocketAddress: &envoy_config_core_v3.SocketAddress{
								Protocol: envoy_config_core_v3.SocketAddress_TCP,
								Address:  ipAddress,
								PortSpecifier: &envoy_config_core_v3.SocketAddress_PortValue{
									PortValue: host.Port,
								},
							},
						},
					},
					HealthCheckConfig: &envoy_config_endpoint_v3.Endpoint_HealthCheckConfig{
						Hostname: ipAddress,
					},
				},
			},
		}

		lbEndpoints = append(lbEndpoints, lbEndpoint)
	}

	return lbEndpoints, nil
}

func resolveIpAddresses(ctx context.Context, address string, resolver DnsResolver) ([]string, error) {
	if resolver == nil {
		return nil, eris.Errorf("Static localized spec defined an address that couldn't be parsed as an IP (%s), "+
			"would have resolved as a hostname but the configured DNS resolver was nil", address)
	}

	ipAddrs, err := resolver.Resolve(ctx, address)
	if err != nil {
		return nil, err
	}

	var ipAddresses []string
	for _, ipAddr := range ipAddrs {
		ipAddresses = append(ipAddresses, ipAddr.String())
	}
	return ipAddresses, nil
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

func getSniAddrForHost(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) string {
	if in.GetSniAddr() != "" {
		return in.GetSniAddr()
	}
	if spec.GetAutoSniRewrite() == nil || spec.GetAutoSniRewrite().GetValue() {
		return in.GetAddr()
	}
	return ""
}

func getMetadataForHost(spec *v1static_localized.UpstreamSpec, in *v1static_localized.Host) *envoy_config_core_v3.Metadata {
	if in == nil {
		return nil
	}
	var meta *envoy_config_core_v3.Metadata
	sniaddr := getSniAddrForHost(spec, in)
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
	return fmt.Sprintf("%s;%s:%d", getSniAddrForHost(spec, in), in.Addr, in.Port)
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

func getTransportSocket(sniHostname string) *envoy_config_core_v3.TransportSocket {
	// TODO: support client certificates
	tlsContext := &envoyauth.UpstreamTlsContext{
		// TODO(yuval-k): Add verification context
		Sni: sniHostname,
	}
	return &envoy_config_core_v3.TransportSocket{
		Name: wellknown.TransportSocketTls,
		ConfigType: &envoy_config_core_v3.TransportSocket_TypedConfig{
			TypedConfig: utils.MustMessageToAny(tlsContext),
		},
	}
}

func getTransportSocketMatches(spec *v1static_localized.UpstreamSpec, transportSocket *envoy_config_core_v3.TransportSocket) ([]*envoy_config_cluster_v3.Cluster_TransportSocketMatch, error) {
	var transportSocketMatches []*envoy_config_cluster_v3.Cluster_TransportSocketMatch

	for _, localizedHosts := range spec.LocalizedHosts {
		for _, host := range localizedHosts.Hosts {
			sniAddress := getSniAddrForHost(spec, host)
			if sniAddress == "" {
				continue
			}
			ts, err := mutateSni(transportSocket, sniAddress)
			if err != nil {
				return nil, err
			}
			transportSocketMatches = append(transportSocketMatches, &envoy_config_cluster_v3.Cluster_TransportSocketMatch{
				Name:            name(spec, host),
				Match:           metadataMatch(spec, host),
				TransportSocket: ts,
			})
		}
	}
	return transportSocketMatches, nil
}
