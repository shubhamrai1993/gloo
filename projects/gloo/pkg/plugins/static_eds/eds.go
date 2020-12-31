package static_eds

import (
	"context"
	"net"
	"sync"
	"time"

	v1static_eds "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static_eds"

	"github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/go-utils/hashutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

func (p *plugin) WatchEndpoints(writeNamespace string, upstreamsToTrack v1.UpstreamList, opts clients.WatchOpts) (<-chan v1.EndpointList, <-chan error, error) {

	var previousHash uint64

	// TODO - // Filter out non-staticEds upstreams

	errChan := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()

	endpointsChan := make(chan v1.EndpointList)
	wg.Add(1)
	go func() {
		defer close(endpointsChan)
		defer wg.Done()

		// Create a new context for each loop, cancel it before each loop
		var cancel context.CancelFunc = func() {}
		// Use closure to allow cancel function to be updated as context changes
		defer func() { cancel() }()

		timer := time.NewTicker(DefaultDnsPollingInterval)

		publishEndpoints := func(endpointListByLocality map[*v1static_eds.Locality]v1.EndpointList) bool {
			if opts.Ctx.Err() != nil {
				return false
			}

			// EDS publishes an aggregated list of Endpoints
			var endpoints v1.EndpointList
			for _, endpointList := range endpointListByLocality {
				endpoints = append(endpoints, endpointList...)
			}

			select {
			case <-opts.Ctx.Done():
				return false
			case endpointsChan <- endpoints:
				p.endpointListByLocality = endpointListByLocality
			}
			return true
		}

		for {
			// don't leak the timer.
			defer timer.Stop()
			select {
			case <-timer.C:
				// Poll to ensure any DNS updates get picked up in endpoints for EDS
				endpointsByLocality := buildEndpointListByLocality()

				currentHash := hashutils.MustHash(endpointsByLocality)
				if previousHash == currentHash {
					continue
				}

				previousHash = currentHash

				if !publishEndpoints(endpointsByLocality) {
					return
				}

			case <-opts.Ctx.Done():
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()
	return endpointsChan, errChan, nil
}

func buildEndpointListByLocality() map[*v1static_eds.Locality]v1.EndpointList {
	// TODO
	return nil
}

func resolveIpAddresses(ctx context.Context, address string, resolver DnsResolver) ([]string, error) {
	addr := net.ParseIP(address)
	if addr != nil {
		// the address is an IP address, no need to resolve it
		return []string{address}, nil
	}

	if resolver == nil {
		return nil, eris.Errorf("Static eds spec defined an address that couldn't be parsed as an IP (%s), "+
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
