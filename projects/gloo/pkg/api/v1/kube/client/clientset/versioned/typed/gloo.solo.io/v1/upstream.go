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

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/apis/gloo.solo.io/v1"
	scheme "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/kube/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UpstreamsGetter has a method to return a UpstreamInterface.
// A group's client should implement this interface.
type UpstreamsGetter interface {
	Upstreams(namespace string) UpstreamInterface
}

// UpstreamInterface has methods to work with Upstream resources.
type UpstreamInterface interface {
	Create(*v1.Upstream) (*v1.Upstream, error)
	Update(*v1.Upstream) (*v1.Upstream, error)
	UpdateStatus(*v1.Upstream) (*v1.Upstream, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Upstream, error)
	List(opts metav1.ListOptions) (*v1.UpstreamList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Upstream, err error)
	UpstreamExpansion
}

// upstreams implements UpstreamInterface
type upstreams struct {
	client rest.Interface
	ns     string
}

// newUpstreams returns a Upstreams
func newUpstreams(c *GlooV1Client, namespace string) *upstreams {
	return &upstreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the upstream, and returns the corresponding upstream object, and an error if there is any.
func (c *upstreams) Get(name string, options metav1.GetOptions) (result *v1.Upstream, err error) {
	result = &v1.Upstream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Upstreams that match those selectors.
func (c *upstreams) List(opts metav1.ListOptions) (result *v1.UpstreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UpstreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("upstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested upstreams.
func (c *upstreams) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("upstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a upstream and creates it.  Returns the server's representation of the upstream, and an error, if there is any.
func (c *upstreams) Create(upstream *v1.Upstream) (result *v1.Upstream, err error) {
	result = &v1.Upstream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("upstreams").
		Body(upstream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a upstream and updates it. Returns the server's representation of the upstream, and an error, if there is any.
func (c *upstreams) Update(upstream *v1.Upstream) (result *v1.Upstream, err error) {
	result = &v1.Upstream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upstreams").
		Name(upstream.Name).
		Body(upstream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *upstreams) UpdateStatus(upstream *v1.Upstream) (result *v1.Upstream, err error) {
	result = &v1.Upstream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("upstreams").
		Name(upstream.Name).
		SubResource("status").
		Body(upstream).
		Do().
		Into(result)
	return
}

// Delete takes name of the upstream and deletes it. Returns an error if one occurs.
func (c *upstreams) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *upstreams) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("upstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched upstream.
func (c *upstreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Upstream, err error) {
	result = &v1.Upstream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("upstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
