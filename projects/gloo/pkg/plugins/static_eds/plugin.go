package static_eds

import (
	"net/url"
	"time"

	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/gloo/projects/gloo/pkg/xds"

	"github.com/rotisserie/eris"

	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_endpoint_v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	v1static_eds "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static_eds"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var _ plugins.Plugin = new(plugin)
var _ discovery.DiscoveryPlugin = new(plugin)
var _ plugins.UpstreamPlugin = new(plugin)
var _ plugins.EndpointPlugin = new(plugin)

var (
	DefaultDnsPollingInterval = 5 * time.Second
)

type plugin struct {
	resolver DnsResolver
	settings *v1.Settings

	endpointListByLocality map[*v1static_eds.Locality]v1.EndpointList
}

func NewPlugin(resolver DnsResolver) plugins.Plugin {
	return &plugin{
		resolver: resolver,
	}
}

func (p *plugin) Resolve(u *v1.Upstream) (*url.URL, error) {
	// TODO
	return nil, nil
}

func (p *plugin) Init(params plugins.InitParams) error {
	p.settings = params.Settings
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoy_config_cluster_v3.Cluster) error {
	if _, ok := in.UpstreamType.(*v1.Upstream_StaticEds); !ok {
		// not ours
		return nil
	}

	// configure the cluster to use EDS:ADS
	xds.SetEdsOnCluster(out, p.settings)

	return nil
}

func (p *plugin) ProcessEndpoints(params plugins.Params, in *v1.Upstream, out *envoy_config_endpoint_v3.ClusterLoadAssignment) error {
	upstreamConfig := in.GetStaticEds()
	if upstreamConfig == nil {
		return nil
	}

	// A ClusterLoadAssignment contains multiple groupings of endpoints by locality
	// At the moment however, our EDS implementation does not support locality, and therefore we create clusters
	// where all endpoints are grouped together.
	//
	// This upstream enables users to define a ClusterLoadAssignment with multiple localities and load balance weights
	// Since EDS does not support this feature (yet), we need to do the following:
	// 	1. During EDS we persist each EndpointList by it's associated Locality
	//	2. During ProcessEndpoints, our ClusterLoadAssignment has a list of LbEndpoints all grouped together in a
	//		single LocalityLbEndpoints instance. To group these LbEndpoints by Locality, we:
	//			2a:	Extract the full list of LbEndpoints from the ClusterLoadAssignment
	//			2b: Map each LbEndpoint to it's associated Locality from Step 1.
	//			2c: Append relevant ClusterLoadAssignment properties defined on Upstream
	//			2d: Overwrite ClusterLoadAssignment with endpoints separated by Locality

	if len(out.Endpoints) == 0 {
		// nothing to do
		return nil
	}

	// Step 2a
	clusterLbEndpoints, err := getLbEndpointsFromClusterLoadAssignment(out)
	if err != nil {
		return err
	}

	// Step 2b
	lbEndpointsByLocality := groupLbEndpointsByLocality(clusterLbEndpoints, p.endpointListByLocality)

	// Step 2c
	localityLbEndpoints := buildLocalityLbEndpoints(upstreamConfig, lbEndpointsByLocality)

	// Step 2d
	out.Endpoints = localityLbEndpoints

	return nil
}

func getLbEndpointsFromClusterLoadAssignment(cla *envoy_config_endpoint_v3.ClusterLoadAssignment) ([]*envoy_config_endpoint_v3.LbEndpoint, error) {
	if len(cla.Endpoints) > 1 {
		// The implementation of this Upstream plugin is based on the expectation that we group all endpoints
		// into a single LocalityLbEndpoints instance. If this condition is met, it means our assumption
		// is no longer correct, and we likely need to update this Upstreams implementation

		return nil, eris.New("TODO - create error")
	}

	return cla.Endpoints[0].GetLbEndpoints(), nil
}

func groupLbEndpointsByLocality(
	lbEndpoints []*envoy_config_endpoint_v3.LbEndpoint,
	endpointListGrouping map[*v1static_eds.Locality]v1.EndpointList,
) map[*v1static_eds.Locality][]*envoy_config_endpoint_v3.LbEndpoint {
	return nil
}

func buildLocalityLbEndpoints(
	upstreamSpec *v1static_eds.UpstreamSpec,
	lbEndpointsByLocality map[*v1static_eds.Locality][]*envoy_config_endpoint_v3.LbEndpoint,
) []*envoy_config_endpoint_v3.LocalityLbEndpoints {
	return nil
}
