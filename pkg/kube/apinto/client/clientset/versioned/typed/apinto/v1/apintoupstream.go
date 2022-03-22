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
	"context"
	"time"

	scheme "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/clientset/versioned/scheme"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApintoUpstreamsGetter has a method to return a ApintoUpstreamInterface.
// A group's client should implement this interface.
type ApintoUpstreamsGetter interface {
	ApintoUpstreams(namespace string) ApintoUpstreamInterface
}

// ApintoUpstreamInterface has methods to work with ApintoUpstream resources.
type ApintoUpstreamInterface interface {
	Create(ctx context.Context, apintoUpstream *v1.ApintoUpstream, opts metav1.CreateOptions) (*v1.ApintoUpstream, error)
	Update(ctx context.Context, apintoUpstream *v1.ApintoUpstream, opts metav1.UpdateOptions) (*v1.ApintoUpstream, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ApintoUpstream, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ApintoUpstreamList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApintoUpstream, err error)
	ApintoUpstreamExpansion
}

// apintoUpstreams implements ApintoUpstreamInterface
type apintoUpstreams struct {
	client rest.Interface
	ns     string
}

// newApintoUpstreams returns a ApintoUpstreams
func newApintoUpstreams(c *ApintoV1Client, namespace string) *apintoUpstreams {
	return &apintoUpstreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apintoUpstream, and returns the corresponding apintoUpstream object, and an error if there is any.
func (c *apintoUpstreams) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ApintoUpstream, err error) {
	result = &v1.ApintoUpstream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apintoupstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApintoUpstreams that match those selectors.
func (c *apintoUpstreams) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ApintoUpstreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApintoUpstreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apintoupstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apintoUpstreams.
func (c *apintoUpstreams) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apintoupstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apintoUpstream and creates it.  Returns the server's representation of the apintoUpstream, and an error, if there is any.
func (c *apintoUpstreams) Create(ctx context.Context, apintoUpstream *v1.ApintoUpstream, opts metav1.CreateOptions) (result *v1.ApintoUpstream, err error) {
	result = &v1.ApintoUpstream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apintoupstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apintoUpstream).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apintoUpstream and updates it. Returns the server's representation of the apintoUpstream, and an error, if there is any.
func (c *apintoUpstreams) Update(ctx context.Context, apintoUpstream *v1.ApintoUpstream, opts metav1.UpdateOptions) (result *v1.ApintoUpstream, err error) {
	result = &v1.ApintoUpstream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apintoupstreams").
		Name(apintoUpstream.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apintoUpstream).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apintoUpstream and deletes it. Returns an error if one occurs.
func (c *apintoUpstreams) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apintoupstreams").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apintoUpstreams) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apintoupstreams").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apintoUpstream.
func (c *apintoUpstreams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApintoUpstream, err error) {
	result = &v1.ApintoUpstream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apintoupstreams").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
