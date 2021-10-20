/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/Tencent/bk-bcs/bcs-k8s/bcs-general-pod-autoscaler/pkg/apis/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGeneralPodAutoscalers implements GeneralPodAutoscalerInterface
type FakeGeneralPodAutoscalers struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var generalpodautoscalersResource = schema.GroupVersionResource{Group: "autoscaling.tkex.tencent.com", Version: "v1alpha1", Resource: "generalpodautoscalers"}

var generalpodautoscalersKind = schema.GroupVersionKind{Group: "autoscaling.tkex.tencent.com", Version: "v1alpha1", Kind: "GeneralPodAutoscaler"}

// Get takes name of the generalPodAutoscaler, and returns the corresponding generalPodAutoscaler object, and an error if there is any.
func (c *FakeGeneralPodAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(generalpodautoscalersResource, c.ns, name), &v1alpha1.GeneralPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), err
}

// List takes label and field selectors, and returns the list of GeneralPodAutoscalers that match those selectors.
func (c *FakeGeneralPodAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.GeneralPodAutoscalerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(generalpodautoscalersResource, generalpodautoscalersKind, c.ns, opts), &v1alpha1.GeneralPodAutoscalerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GeneralPodAutoscalerList{ListMeta: obj.(*v1alpha1.GeneralPodAutoscalerList).ListMeta}
	for _, item := range obj.(*v1alpha1.GeneralPodAutoscalerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested generalPodAutoscalers.
func (c *FakeGeneralPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(generalpodautoscalersResource, c.ns, opts))

}

// Create takes the representation of a generalPodAutoscaler and creates it.  Returns the server's representation of the generalPodAutoscaler, and an error, if there is any.
func (c *FakeGeneralPodAutoscalers) Create(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(generalpodautoscalersResource, c.ns, generalPodAutoscaler), &v1alpha1.GeneralPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), err
}

// Update takes the representation of a generalPodAutoscaler and updates it. Returns the server's representation of the generalPodAutoscaler, and an error, if there is any.
func (c *FakeGeneralPodAutoscalers) Update(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(generalpodautoscalersResource, c.ns, generalPodAutoscaler), &v1alpha1.GeneralPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGeneralPodAutoscalers) UpdateStatus(generalPodAutoscaler *v1alpha1.GeneralPodAutoscaler) (*v1alpha1.GeneralPodAutoscaler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(generalpodautoscalersResource, "status", c.ns, generalPodAutoscaler), &v1alpha1.GeneralPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), err
}

// Delete takes name of the generalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *FakeGeneralPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(generalpodautoscalersResource, c.ns, name), &v1alpha1.GeneralPodAutoscaler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGeneralPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(generalpodautoscalersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GeneralPodAutoscalerList{})
	return err
}

// Patch applies the patch and returns the patched generalPodAutoscaler.
func (c *FakeGeneralPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GeneralPodAutoscaler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(generalpodautoscalersResource, c.ns, name, pt, data, subresources...), &v1alpha1.GeneralPodAutoscaler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GeneralPodAutoscaler), err
}
