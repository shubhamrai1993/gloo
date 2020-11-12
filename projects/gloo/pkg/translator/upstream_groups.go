package translator

import (
	errors "github.com/rotisserie/eris"
	"github.com/solo-io/gloo-edge/projects/gloo/pkg/plugins"
	usconversions "github.com/solo-io/gloo-edge/projects/gloo/pkg/upstreams"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
)

func (t *translatorInstance) verifyUpstreamGroups(params plugins.Params, reports reporter.ResourceReports) {

	upstreams := params.Snapshot.Upstreams
	upstreamGroups := params.Snapshot.UpstreamGroups

	for _, ug := range upstreamGroups {
		for i, dest := range ug.Destinations {
			if dest.Destination == nil {
				reports.AddError(ug, errors.Errorf("destination # %d: destination is nil", i+1))
				continue
			}

			if upstream := dest.GetDestination().GetUpstream(); upstream != nil && upstream.GetNamespace() == "" {
				parentMetadata := ug.GetMetadata()
				upstream.Namespace = parentMetadata.GetNamespace()
			}

			upRef, err := usconversions.DestinationToUpstreamRef(dest.Destination)
			if err != nil {
				reports.AddError(ug, err)
				continue
			}

			if _, err := upstreams.Find(upRef.Namespace, upRef.Name); err != nil {
				reports.AddError(ug, errors.Wrapf(err, "destination # %d: upstream not found", i+1))
				continue
			}
		}

	}

}
