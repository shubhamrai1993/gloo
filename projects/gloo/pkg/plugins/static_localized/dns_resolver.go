package static_localized

import (
	"context"
	"net"

	"github.com/rotisserie/eris"
)

//go:generate mockgen -destination ./mocks/dnsresolver_mock.go github.com/solo-io/gloo/projects/gloo/pkg/plugins/static_localized DnsResolver
//go:generate gofmt -w ./mocks/
//go:generate goimports -w ./mocks/

type DnsResolver interface {
	Resolve(ctx context.Context, address string) ([]net.IPAddr, error)
}

type StaticLocalizedDnsResolver struct {
}

func (c *StaticLocalizedDnsResolver) Resolve(ctx context.Context, address string) ([]net.IPAddr, error) {
	res := net.Resolver{
		PreferGo: true, // otherwise we may use cgo which doesn't resolve on my mac in testing
		Dial: func(ctx context.Context, network, address string) (conn net.Conn, err error) {
			// DNS typically uses UDP and falls back to TCP if the response size is greater than one packet
			// (originally 512 bytes). we use TCP to ensure we receive all IPs in a large DNS response
			return net.Dial("tcp", address)
		},
	}
	ipAddrs, err := res.LookupIPAddr(ctx, address)
	if err != nil {
		return nil, err
	}
	if len(ipAddrs) == 0 {
		return nil, eris.Errorf("Static localized spec contained a host address that couldn't be parsed as an IP (%s), "+
			"could not resolve as a hostname", address)
	}
	return ipAddrs, nil
}
