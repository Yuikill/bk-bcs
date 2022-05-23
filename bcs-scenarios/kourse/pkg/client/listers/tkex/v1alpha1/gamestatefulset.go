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

package v1alpha1

import (
	v1alpha1 "github.com/bk-bcs/bcs-scenarios/kourse/pkg/apis/tkex/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GameStatefulSetLister helps list GameStatefulSets.
// All objects returned here must be treated as read-only.
type GameStatefulSetLister interface {
	// List lists all GameStatefulSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GameStatefulSet, err error)
	// GameStatefulSets returns an object that can list and get GameStatefulSets.
	GameStatefulSets(namespace string) GameStatefulSetNamespaceLister
	GameStatefulSetListerExpansion
}

// gameStatefulSetLister implements the GameStatefulSetLister interface.
type gameStatefulSetLister struct {
	indexer cache.Indexer
}

// NewGameStatefulSetLister returns a new GameStatefulSetLister.
func NewGameStatefulSetLister(indexer cache.Indexer) GameStatefulSetLister {
	return &gameStatefulSetLister{indexer: indexer}
}

// List lists all GameStatefulSets in the indexer.
func (s *gameStatefulSetLister) List(selector labels.Selector) (ret []*v1alpha1.GameStatefulSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GameStatefulSet))
	})
	return ret, err
}

// GameStatefulSets returns an object that can list and get GameStatefulSets.
func (s *gameStatefulSetLister) GameStatefulSets(namespace string) GameStatefulSetNamespaceLister {
	return gameStatefulSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GameStatefulSetNamespaceLister helps list and get GameStatefulSets.
// All objects returned here must be treated as read-only.
type GameStatefulSetNamespaceLister interface {
	// List lists all GameStatefulSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GameStatefulSet, err error)
	// Get retrieves the GameStatefulSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GameStatefulSet, error)
	GameStatefulSetNamespaceListerExpansion
}

// gameStatefulSetNamespaceLister implements the GameStatefulSetNamespaceLister
// interface.
type gameStatefulSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GameStatefulSets in the indexer for a given namespace.
func (s gameStatefulSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GameStatefulSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GameStatefulSet))
	})
	return ret, err
}

// Get retrieves the GameStatefulSet from the indexer for a given namespace and name.
func (s gameStatefulSetNamespaceLister) Get(name string) (*v1alpha1.GameStatefulSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gamestatefulset"), name)
	}
	return obj.(*v1alpha1.GameStatefulSet), nil
}
