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

// ApintoDiscoveriesGetter has a method to return a ApintoDiscoveryInterface.
// A group's client should implement this interface.
type ApintoDiscoveriesGetter interface {
	ApintoDiscoveries(namespace string) ApintoDiscoveryInterface
}

// ApintoDiscoveryInterface has methods to work with ApintoDiscovery resources.
type ApintoDiscoveryInterface interface {
	Create(ctx context.Context, apintoDiscovery *v1.ApintoDiscovery, opts metav1.CreateOptions) (*v1.ApintoDiscovery, error)
	Update(ctx context.Context, apintoDiscovery *v1.ApintoDiscovery, opts metav1.UpdateOptions) (*v1.ApintoDiscovery, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ApintoDiscovery, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ApintoDiscoveryList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApintoDiscovery, err error)
	ApintoDiscoveryExpansion
}

// apintoDiscoveries implements ApintoDiscoveryInterface
type apintoDiscoveries struct {
	client rest.Interface
	ns     string
}

// newApintoDiscoveries returns a ApintoDiscoveries
func newApintoDiscoveries(c *ApintoV1Client, namespace string) *apintoDiscoveries {
	return &apintoDiscoveries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apintoDiscovery, and returns the corresponding apintoDiscovery object, and an error if there is any.
func (c *apintoDiscoveries) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ApintoDiscovery, err error) {
	result = &v1.ApintoDiscovery{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApintoDiscoveries that match those selectors.
func (c *apintoDiscoveries) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ApintoDiscoveryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApintoDiscoveryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apintoDiscoveries.
func (c *apintoDiscoveries) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a apintoDiscovery and creates it.  Returns the server's representation of the apintoDiscovery, and an error, if there is any.
func (c *apintoDiscoveries) Create(ctx context.Context, apintoDiscovery *v1.ApintoDiscovery, opts metav1.CreateOptions) (result *v1.ApintoDiscovery, err error) {
	result = &v1.ApintoDiscovery{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apintoDiscovery).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a apintoDiscovery and updates it. Returns the server's representation of the apintoDiscovery, and an error, if there is any.
func (c *apintoDiscoveries) Update(ctx context.Context, apintoDiscovery *v1.ApintoDiscovery, opts metav1.UpdateOptions) (result *v1.ApintoDiscovery, err error) {
	result = &v1.ApintoDiscovery{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		Name(apintoDiscovery.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(apintoDiscovery).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the apintoDiscovery and deletes it. Returns an error if one occurs.
func (c *apintoDiscoveries) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apintoDiscoveries) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apintodiscoveries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched apintoDiscovery.
func (c *apintoDiscoveries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApintoDiscovery, err error) {
	result = &v1.ApintoDiscovery{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apintodiscoveries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}