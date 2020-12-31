package static_eds

import (
	"github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

// StaticEds Upstreams are manually defined by the user and therefore do not rely on UDS
// Since the DiscoveryPlugin interface must implement UDS and EDS, this no-op UDS implementation is required

var (
	InvalidSpecTypeError = func(us *v1.Upstream, name string) error {
		return eris.Errorf("internal error: invalid %s spec, "+
			"expected *v1.Upstream_StaticEds, got  %T", name, us)
	}
)

func (p *plugin) DiscoverUpstreams(watchNamespaces []string, writeNamespace string, opts clients.WatchOpts, discOpts discovery.Opts) (chan v1.UpstreamList, chan error, error) {
	return nil, nil, nil
}

func (p *plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	originalSpec, ok := original.UpstreamType.(*v1.Upstream_StaticEds)
	if !ok {
		return false, InvalidSpecTypeError(original, "original")
	}
	desiredSpec, ok := desired.UpstreamType.(*v1.Upstream_StaticEds)
	if !ok {
		return false, InvalidSpecTypeError(desired, "desired")
	}

	return !originalSpec.StaticEds.Equal(desiredSpec.StaticEds), nil
}
