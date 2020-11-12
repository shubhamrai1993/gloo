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

package v1

import (
	"time"

	v1 "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	scheme "github.com/solo-io/gloo-edge/projects/gloo/pkg/api/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProxiesGetter has a method to return a ProxyInterface.
// A group's client should implement this interface.
type ProxiesGetter interface {
	Proxies(namespace string) ProxyInterface
}

// ProxyInterface has methods to work with Proxy resources.
type ProxyInterface interface {
	Create(*v1.Proxy) (*v1.Proxy, error)
	Update(*v1.Proxy) (*v1.Proxy, error)
	UpdateStatus(*v1.Proxy) (*v1.Proxy, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Proxy, error)
	List(opts metav1.ListOptions) (*v1.ProxyList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Proxy, err error)
	ProxyExpansion
}

// proxies implements ProxyInterface
type proxies struct {
	client rest.Interface
	ns     string
}

// newProxies returns a Proxies
func newProxies(c *GlooV1Client, namespace string) *proxies {
	return &proxies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the proxy, and returns the corresponding proxy object, and an error if there is any.
func (c *proxies) Get(name string, options metav1.GetOptions) (result *v1.Proxy, err error) {
	result = &v1.Proxy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Proxies that match those selectors.
func (c *proxies) List(opts metav1.ListOptions) (result *v1.ProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ProxyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested proxies.
func (c *proxies) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("proxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a proxy and creates it.  Returns the server's representation of the proxy, and an error, if there is any.
func (c *proxies) Create(proxy *v1.Proxy) (result *v1.Proxy, err error) {
	result = &v1.Proxy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("proxies").
		Body(proxy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a proxy and updates it. Returns the server's representation of the proxy, and an error, if there is any.
func (c *proxies) Update(proxy *v1.Proxy) (result *v1.Proxy, err error) {
	result = &v1.Proxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxies").
		Name(proxy.Name).
		Body(proxy).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *proxies) UpdateStatus(proxy *v1.Proxy) (result *v1.Proxy, err error) {
	result = &v1.Proxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxies").
		Name(proxy.Name).
		SubResource("status").
		Body(proxy).
		Do().
		Into(result)
	return
}

// Delete takes name of the proxy and deletes it. Returns an error if one occurs.
func (c *proxies) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *proxies) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched proxy.
func (c *proxies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Proxy, err error) {
	result = &v1.Proxy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("proxies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
