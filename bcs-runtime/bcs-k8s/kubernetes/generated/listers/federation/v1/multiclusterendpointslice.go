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
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Tencent/bk-bcs/bcs-runtime/bcs-k8s/kubernetes/apis/federation/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MultiClusterEndpointSliceLister helps list MultiClusterEndpointSlices.
// All objects returned here must be treated as read-only.
type MultiClusterEndpointSliceLister interface {
	// List lists all MultiClusterEndpointSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.MultiClusterEndpointSlice, err error)
	// MultiClusterEndpointSlices returns an object that can list and get MultiClusterEndpointSlices.
	MultiClusterEndpointSlices(namespace string) MultiClusterEndpointSliceNamespaceLister
	MultiClusterEndpointSliceListerExpansion
}

// multiClusterEndpointSliceLister implements the MultiClusterEndpointSliceLister interface.
type multiClusterEndpointSliceLister struct {
	indexer cache.Indexer
}

// NewMultiClusterEndpointSliceLister returns a new MultiClusterEndpointSliceLister.
func NewMultiClusterEndpointSliceLister(indexer cache.Indexer) MultiClusterEndpointSliceLister {
	return &multiClusterEndpointSliceLister{indexer: indexer}
}

// List lists all MultiClusterEndpointSlices in the indexer.
func (s *multiClusterEndpointSliceLister) List(selector labels.Selector) (ret []*v1.MultiClusterEndpointSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MultiClusterEndpointSlice))
	})
	return ret, err
}

// MultiClusterEndpointSlices returns an object that can list and get MultiClusterEndpointSlices.
func (s *multiClusterEndpointSliceLister) MultiClusterEndpointSlices(namespace string) MultiClusterEndpointSliceNamespaceLister {
	return multiClusterEndpointSliceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MultiClusterEndpointSliceNamespaceLister helps list and get MultiClusterEndpointSlices.
// All objects returned here must be treated as read-only.
type MultiClusterEndpointSliceNamespaceLister interface {
	// List lists all MultiClusterEndpointSlices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.MultiClusterEndpointSlice, err error)
	// Get retrieves the MultiClusterEndpointSlice from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.MultiClusterEndpointSlice, error)
	MultiClusterEndpointSliceNamespaceListerExpansion
}

// multiClusterEndpointSliceNamespaceLister implements the MultiClusterEndpointSliceNamespaceLister
// interface.
type multiClusterEndpointSliceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MultiClusterEndpointSlices in the indexer for a given namespace.
func (s multiClusterEndpointSliceNamespaceLister) List(selector labels.Selector) (ret []*v1.MultiClusterEndpointSlice, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.MultiClusterEndpointSlice))
	})
	return ret, err
}

// Get retrieves the MultiClusterEndpointSlice from the indexer for a given namespace and name.
func (s multiClusterEndpointSliceNamespaceLister) Get(name string) (*v1.MultiClusterEndpointSlice, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("multiclusterendpointslice"), name)
	}
	return obj.(*v1.MultiClusterEndpointSlice), nil
}
