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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApintoRouterLister helps list ApintoRouters.
// All objects returned here must be treated as read-only.
type ApintoRouterLister interface {
	// List lists all ApintoRouters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApintoRouter, err error)
	// ApintoRouters returns an object that can list and get ApintoRouters.
	ApintoRouters(namespace string) ApintoRouterNamespaceLister
	ApintoRouterListerExpansion
}

// apintoRouterLister implements the ApintoRouterLister interface.
type apintoRouterLister struct {
	indexer cache.Indexer
}

// NewApintoRouterLister returns a new ApintoRouterLister.
func NewApintoRouterLister(indexer cache.Indexer) ApintoRouterLister {
	return &apintoRouterLister{indexer: indexer}
}

// List lists all ApintoRouters in the indexer.
func (s *apintoRouterLister) List(selector labels.Selector) (ret []*v1.ApintoRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApintoRouter))
	})
	return ret, err
}

// ApintoRouters returns an object that can list and get ApintoRouters.
func (s *apintoRouterLister) ApintoRouters(namespace string) ApintoRouterNamespaceLister {
	return apintoRouterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApintoRouterNamespaceLister helps list and get ApintoRouters.
// All objects returned here must be treated as read-only.
type ApintoRouterNamespaceLister interface {
	// List lists all ApintoRouters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApintoRouter, err error)
	// Get retrieves the ApintoRouter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ApintoRouter, error)
	ApintoRouterNamespaceListerExpansion
}

// apintoRouterNamespaceLister implements the ApintoRouterNamespaceLister
// interface.
type apintoRouterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApintoRouters in the indexer for a given namespace.
func (s apintoRouterNamespaceLister) List(selector labels.Selector) (ret []*v1.ApintoRouter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApintoRouter))
	})
	return ret, err
}

// Get retrieves the ApintoRouter from the indexer for a given namespace and name.
func (s apintoRouterNamespaceLister) Get(name string) (*v1.ApintoRouter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("apintorouter"), name)
	}
	return obj.(*v1.ApintoRouter), nil
}
