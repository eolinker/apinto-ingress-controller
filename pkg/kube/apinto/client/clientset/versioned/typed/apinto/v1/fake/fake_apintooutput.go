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
	"context"

	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApintoOutputs implements ApintoOutputInterface
type FakeApintoOutputs struct {
	Fake *FakeApintoV1
	ns   string
}

var apintooutputsResource = schema.GroupVersionResource{Group: "apinto.com", Version: "v1", Resource: "apintooutputs"}

var apintooutputsKind = schema.GroupVersionKind{Group: "apinto.com", Version: "v1", Kind: "ApintoOutput"}

// Get takes name of the apintoOutput, and returns the corresponding apintoOutput object, and an error if there is any.
func (c *FakeApintoOutputs) Get(ctx context.Context, name string, options v1.GetOptions) (result *apintov1.ApintoOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apintooutputsResource, c.ns, name), &apintov1.ApintoOutput{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoOutput), err
}

// List takes label and field selectors, and returns the list of ApintoOutputs that match those selectors.
func (c *FakeApintoOutputs) List(ctx context.Context, opts v1.ListOptions) (result *apintov1.ApintoOutputList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apintooutputsResource, apintooutputsKind, c.ns, opts), &apintov1.ApintoOutputList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &apintov1.ApintoOutputList{ListMeta: obj.(*apintov1.ApintoOutputList).ListMeta}
	for _, item := range obj.(*apintov1.ApintoOutputList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apintoOutputs.
func (c *FakeApintoOutputs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apintooutputsResource, c.ns, opts))

}

// Create takes the representation of a apintoOutput and creates it.  Returns the server's representation of the apintoOutput, and an error, if there is any.
func (c *FakeApintoOutputs) Create(ctx context.Context, apintoOutput *apintov1.ApintoOutput, opts v1.CreateOptions) (result *apintov1.ApintoOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apintooutputsResource, c.ns, apintoOutput), &apintov1.ApintoOutput{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoOutput), err
}

// Update takes the representation of a apintoOutput and updates it. Returns the server's representation of the apintoOutput, and an error, if there is any.
func (c *FakeApintoOutputs) Update(ctx context.Context, apintoOutput *apintov1.ApintoOutput, opts v1.UpdateOptions) (result *apintov1.ApintoOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apintooutputsResource, c.ns, apintoOutput), &apintov1.ApintoOutput{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoOutput), err
}

// Delete takes name of the apintoOutput and deletes it. Returns an error if one occurs.
func (c *FakeApintoOutputs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(apintooutputsResource, c.ns, name, opts), &apintov1.ApintoOutput{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApintoOutputs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apintooutputsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &apintov1.ApintoOutputList{})
	return err
}

// Patch applies the patch and returns the patched apintoOutput.
func (c *FakeApintoOutputs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *apintov1.ApintoOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apintooutputsResource, c.ns, name, pt, data, subresources...), &apintov1.ApintoOutput{})

	if obj == nil {
		return nil, err
	}
	return obj.(*apintov1.ApintoOutput), err
}