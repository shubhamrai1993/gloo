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

// SettingsesGetter has a method to return a SettingsInterface.
// A group's client should implement this interface.
type SettingsesGetter interface {
	Settingses(namespace string) SettingsInterface
}

// SettingsInterface has methods to work with Settings resources.
type SettingsInterface interface {
	Create(*v1.Settings) (*v1.Settings, error)
	Update(*v1.Settings) (*v1.Settings, error)
	UpdateStatus(*v1.Settings) (*v1.Settings, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Settings, error)
	List(opts metav1.ListOptions) (*v1.SettingsList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Settings, err error)
	SettingsExpansion
}

// settingses implements SettingsInterface
type settingses struct {
	client rest.Interface
	ns     string
}

// newSettingses returns a Settingses
func newSettingses(c *GlooV1Client, namespace string) *settingses {
	return &settingses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the settings, and returns the corresponding settings object, and an error if there is any.
func (c *settingses) Get(name string, options metav1.GetOptions) (result *v1.Settings, err error) {
	result = &v1.Settings{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("settings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Settingses that match those selectors.
func (c *settingses) List(opts metav1.ListOptions) (result *v1.SettingsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SettingsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("settings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested settingses.
func (c *settingses) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("settings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a settings and creates it.  Returns the server's representation of the settings, and an error, if there is any.
func (c *settingses) Create(settings *v1.Settings) (result *v1.Settings, err error) {
	result = &v1.Settings{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("settings").
		Body(settings).
		Do().
		Into(result)
	return
}

// Update takes the representation of a settings and updates it. Returns the server's representation of the settings, and an error, if there is any.
func (c *settingses) Update(settings *v1.Settings) (result *v1.Settings, err error) {
	result = &v1.Settings{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("settings").
		Name(settings.Name).
		Body(settings).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *settingses) UpdateStatus(settings *v1.Settings) (result *v1.Settings, err error) {
	result = &v1.Settings{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("settings").
		Name(settings.Name).
		SubResource("status").
		Body(settings).
		Do().
		Into(result)
	return
}

// Delete takes name of the settings and deletes it. Returns an error if one occurs.
func (c *settingses) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("settings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *settingses) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("settings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched settings.
func (c *settingses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Settings, err error) {
	result = &v1.Settings{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("settings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
