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
	gloosoloiov1 "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProxies implements ProxyInterface
type FakeProxies struct {
	Fake *FakeGlooV1
	ns   string
}

var proxiesResource = schema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "proxies"}

var proxiesKind = schema.GroupVersionKind{Group: "gloo.solo.io", Version: "v1", Kind: "Proxy"}

// Get takes name of the proxy, and returns the corresponding proxy object, and an error if there is any.
func (c *FakeProxies) Get(name string, options v1.GetOptions) (result *gloosoloiov1.Proxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(proxiesResource, c.ns, name), &gloosoloiov1.Proxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloosoloiov1.Proxy), err
}

// List takes label and field selectors, and returns the list of Proxies that match those selectors.
func (c *FakeProxies) List(opts v1.ListOptions) (result *gloosoloiov1.ProxyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(proxiesResource, proxiesKind, c.ns, opts), &gloosoloiov1.ProxyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &gloosoloiov1.ProxyList{ListMeta: obj.(*gloosoloiov1.ProxyList).ListMeta}
	for _, item := range obj.(*gloosoloiov1.ProxyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested proxies.
func (c *FakeProxies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(proxiesResource, c.ns, opts))

}

// Create takes the representation of a proxy and creates it.  Returns the server's representation of the proxy, and an error, if there is any.
func (c *FakeProxies) Create(proxy *gloosoloiov1.Proxy) (result *gloosoloiov1.Proxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(proxiesResource, c.ns, proxy), &gloosoloiov1.Proxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloosoloiov1.Proxy), err
}

// Update takes the representation of a proxy and updates it. Returns the server's representation of the proxy, and an error, if there is any.
func (c *FakeProxies) Update(proxy *gloosoloiov1.Proxy) (result *gloosoloiov1.Proxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(proxiesResource, c.ns, proxy), &gloosoloiov1.Proxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloosoloiov1.Proxy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProxies) UpdateStatus(proxy *gloosoloiov1.Proxy) (*gloosoloiov1.Proxy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(proxiesResource, "status", c.ns, proxy), &gloosoloiov1.Proxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloosoloiov1.Proxy), err
}

// Delete takes name of the proxy and deletes it. Returns an error if one occurs.
func (c *FakeProxies) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(proxiesResource, c.ns, name), &gloosoloiov1.Proxy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProxies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(proxiesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &gloosoloiov1.ProxyList{})
	return err
}

// Patch applies the patch and returns the patched proxy.
func (c *FakeProxies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *gloosoloiov1.Proxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(proxiesResource, c.ns, name, pt, data, subresources...), &gloosoloiov1.Proxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*gloosoloiov1.Proxy), err
}
