/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/kube/client/clientset/versioned/typed/gloo.solo.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeGlooV1 struct {
	*testing.Fake
}

func (c *FakeGlooV1) Artifacts(namespace string) v1.ArtifactInterface {
	return &FakeArtifacts{c, namespace}
}

func (c *FakeGlooV1) Endpoints(namespace string) v1.EndpointInterface {
	return &FakeEndpoints{c, namespace}
}

func (c *FakeGlooV1) Proxies(namespace string) v1.ProxyInterface {
	return &FakeProxies{c, namespace}
}

func (c *FakeGlooV1) Secrets(namespace string) v1.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeGlooV1) Settingses(namespace string) v1.SettingsInterface {
	return &FakeSettingses{c, namespace}
}

func (c *FakeGlooV1) Upstreams(namespace string) v1.UpstreamInterface {
	return &FakeUpstreams{c, namespace}
}

func (c *FakeGlooV1) UpstreamGroups(namespace string) v1.UpstreamGroupInterface {
	return &FakeUpstreamGroups{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGlooV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
