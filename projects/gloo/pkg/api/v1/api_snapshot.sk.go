// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"fmt"
	"hash"
	"hash/fnv"
	"log"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/external/solo/ratelimit"
	enterprise_gloo_solo_io "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"

	"github.com/rotisserie/eris"
	"github.com/solo-io/go-utils/hashutils"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Artifacts        ArtifactList
	Endpoints        EndpointList
	Proxies          ProxyList
	UpstreamGroups   UpstreamGroupList
	Secrets          SecretList
	Upstreams        UpstreamList
	AuthConfigs      enterprise_gloo_solo_io.AuthConfigList
	Ratelimitconfigs github_com_solo_io_gloo_projects_gloo_pkg_api_external_solo_ratelimit.RateLimitConfigList
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Artifacts:        s.Artifacts.Clone(),
		Endpoints:        s.Endpoints.Clone(),
		Proxies:          s.Proxies.Clone(),
		UpstreamGroups:   s.UpstreamGroups.Clone(),
		Secrets:          s.Secrets.Clone(),
		Upstreams:        s.Upstreams.Clone(),
		AuthConfigs:      s.AuthConfigs.Clone(),
		Ratelimitconfigs: s.Ratelimitconfigs.Clone(),
	}
}

func (s ApiSnapshot) Hash(hasher hash.Hash64) (uint64, error) {
	if hasher == nil {
		hasher = fnv.New64()
	}
	if _, err := s.hashArtifacts(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashEndpoints(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashProxies(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreamGroups(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashSecrets(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashUpstreams(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashAuthConfigs(hasher); err != nil {
		return 0, err
	}
	if _, err := s.hashRatelimitconfigs(hasher); err != nil {
		return 0, err
	}
	return hasher.Sum64(), nil
}

func (s ApiSnapshot) hashArtifacts(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Artifacts.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashEndpoints(hasher hash.Hash64) (uint64, error) {
	clonedList := s.Endpoints.Clone()
	for _, v := range clonedList {
		v.Metadata.Annotations = nil
	}
	return hashutils.HashAllSafe(hasher, clonedList.AsInterfaces()...)
}

func (s ApiSnapshot) hashProxies(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Proxies.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreamGroups(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.UpstreamGroups.AsInterfaces()...)
}

func (s ApiSnapshot) hashSecrets(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Secrets.AsInterfaces()...)
}

func (s ApiSnapshot) hashUpstreams(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Upstreams.AsInterfaces()...)
}

func (s ApiSnapshot) hashAuthConfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.AuthConfigs.AsInterfaces()...)
}

func (s ApiSnapshot) hashRatelimitconfigs(hasher hash.Hash64) (uint64, error) {
	return hashutils.HashAllSafe(hasher, s.Ratelimitconfigs.AsInterfaces()...)
}

func (s ApiSnapshot) HashFields() []zap.Field {
	var fields []zap.Field
	hasher := fnv.New64()
	ArtifactsHash, err := s.hashArtifacts(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("artifacts", ArtifactsHash))
	EndpointsHash, err := s.hashEndpoints(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("endpoints", EndpointsHash))
	ProxiesHash, err := s.hashProxies(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("proxies", ProxiesHash))
	UpstreamGroupsHash, err := s.hashUpstreamGroups(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreamGroups", UpstreamGroupsHash))
	SecretsHash, err := s.hashSecrets(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("secrets", SecretsHash))
	UpstreamsHash, err := s.hashUpstreams(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("upstreams", UpstreamsHash))
	AuthConfigsHash, err := s.hashAuthConfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("authConfigs", AuthConfigsHash))
	RatelimitconfigsHash, err := s.hashRatelimitconfigs(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	fields = append(fields, zap.Uint64("ratelimitconfigs", RatelimitconfigsHash))
	snapshotHash, err := s.Hash(hasher)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return append(fields, zap.Uint64("snapshotHash", snapshotHash))
}

type ApiSnapshotStringer struct {
	Version          uint64
	Artifacts        []string
	Endpoints        []string
	Proxies          []string
	UpstreamGroups   []string
	Secrets          []string
	Upstreams        []string
	AuthConfigs      []string
	Ratelimitconfigs []string
}

func (ss ApiSnapshotStringer) String() string {
	s := fmt.Sprintf("ApiSnapshot %v\n", ss.Version)

	s += fmt.Sprintf("  Artifacts %v\n", len(ss.Artifacts))
	for _, name := range ss.Artifacts {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Endpoints %v\n", len(ss.Endpoints))
	for _, name := range ss.Endpoints {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Proxies %v\n", len(ss.Proxies))
	for _, name := range ss.Proxies {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  UpstreamGroups %v\n", len(ss.UpstreamGroups))
	for _, name := range ss.UpstreamGroups {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Secrets %v\n", len(ss.Secrets))
	for _, name := range ss.Secrets {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Upstreams %v\n", len(ss.Upstreams))
	for _, name := range ss.Upstreams {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  AuthConfigs %v\n", len(ss.AuthConfigs))
	for _, name := range ss.AuthConfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	s += fmt.Sprintf("  Ratelimitconfigs %v\n", len(ss.Ratelimitconfigs))
	for _, name := range ss.Ratelimitconfigs {
		s += fmt.Sprintf("    %v\n", name)
	}

	return s
}

func (s ApiSnapshot) Stringer() ApiSnapshotStringer {
	snapshotHash, err := s.Hash(nil)
	if err != nil {
		log.Println(eris.Wrapf(err, "error hashing, this should never happen"))
	}
	return ApiSnapshotStringer{
		Version:          snapshotHash,
		Artifacts:        s.Artifacts.NamespacesDotNames(),
		Endpoints:        s.Endpoints.NamespacesDotNames(),
		Proxies:          s.Proxies.NamespacesDotNames(),
		UpstreamGroups:   s.UpstreamGroups.NamespacesDotNames(),
		Secrets:          s.Secrets.NamespacesDotNames(),
		Upstreams:        s.Upstreams.NamespacesDotNames(),
		AuthConfigs:      s.AuthConfigs.NamespacesDotNames(),
		Ratelimitconfigs: s.Ratelimitconfigs.NamespacesDotNames(),
	}
}
