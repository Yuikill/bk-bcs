/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	federationv1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/federation/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterEndpointSlices implements MultiClusterEndpointSliceInterface
type FakeMultiClusterEndpointSlices struct {
	Fake *FakeFederationV1
	ns   string
}

var multiclusterendpointslicesResource = schema.GroupVersionResource{Group: "federation", Version: "v1", Resource: "multiclusterendpointslices"}

var multiclusterendpointslicesKind = schema.GroupVersionKind{Group: "federation", Version: "v1", Kind: "MultiClusterEndpointSlice"}

// Get takes name of the multiClusterEndpointSlice, and returns the corresponding multiClusterEndpointSlice object, and an error if there is any.
func (c *FakeMultiClusterEndpointSlices) Get(ctx context.Context, name string, options v1.GetOptions) (result *federationv1.MultiClusterEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterendpointslicesResource, c.ns, name), &federationv1.MultiClusterEndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.MultiClusterEndpointSlice), err
}

// List takes label and field selectors, and returns the list of MultiClusterEndpointSlices that match those selectors.
func (c *FakeMultiClusterEndpointSlices) List(ctx context.Context, opts v1.ListOptions) (result *federationv1.MultiClusterEndpointSliceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterendpointslicesResource, multiclusterendpointslicesKind, c.ns, opts), &federationv1.MultiClusterEndpointSliceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &federationv1.MultiClusterEndpointSliceList{ListMeta: obj.(*federationv1.MultiClusterEndpointSliceList).ListMeta}
	for _, item := range obj.(*federationv1.MultiClusterEndpointSliceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterEndpointSlices.
func (c *FakeMultiClusterEndpointSlices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterendpointslicesResource, c.ns, opts))

}

// Create takes the representation of a multiClusterEndpointSlice and creates it.  Returns the server's representation of the multiClusterEndpointSlice, and an error, if there is any.
func (c *FakeMultiClusterEndpointSlices) Create(ctx context.Context, multiClusterEndpointSlice *federationv1.MultiClusterEndpointSlice, opts v1.CreateOptions) (result *federationv1.MultiClusterEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterendpointslicesResource, c.ns, multiClusterEndpointSlice), &federationv1.MultiClusterEndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.MultiClusterEndpointSlice), err
}

// Update takes the representation of a multiClusterEndpointSlice and updates it. Returns the server's representation of the multiClusterEndpointSlice, and an error, if there is any.
func (c *FakeMultiClusterEndpointSlices) Update(ctx context.Context, multiClusterEndpointSlice *federationv1.MultiClusterEndpointSlice, opts v1.UpdateOptions) (result *federationv1.MultiClusterEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterendpointslicesResource, c.ns, multiClusterEndpointSlice), &federationv1.MultiClusterEndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.MultiClusterEndpointSlice), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterEndpointSlices) UpdateStatus(ctx context.Context, multiClusterEndpointSlice *federationv1.MultiClusterEndpointSlice, opts v1.UpdateOptions) (*federationv1.MultiClusterEndpointSlice, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusterendpointslicesResource, "status", c.ns, multiClusterEndpointSlice), &federationv1.MultiClusterEndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.MultiClusterEndpointSlice), err
}

// Delete takes name of the multiClusterEndpointSlice and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterEndpointSlices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiclusterendpointslicesResource, c.ns, name), &federationv1.MultiClusterEndpointSlice{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterEndpointSlices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterendpointslicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &federationv1.MultiClusterEndpointSliceList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterEndpointSlice.
func (c *FakeMultiClusterEndpointSlices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *federationv1.MultiClusterEndpointSlice, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterendpointslicesResource, c.ns, name, pt, data, subresources...), &federationv1.MultiClusterEndpointSlice{})

	if obj == nil {
		return nil, err
	}
	return obj.(*federationv1.MultiClusterEndpointSlice), err
}
