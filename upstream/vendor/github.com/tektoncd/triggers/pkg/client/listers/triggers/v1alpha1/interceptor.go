/*
Copyright 2019 The Tekton Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InterceptorLister helps list Interceptors.
// All objects returned here must be treated as read-only.
type InterceptorLister interface {
	// List lists all Interceptors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Interceptor, err error)
	// Interceptors returns an object that can list and get Interceptors.
	Interceptors(namespace string) InterceptorNamespaceLister
	InterceptorListerExpansion
}

// interceptorLister implements the InterceptorLister interface.
type interceptorLister struct {
	indexer cache.Indexer
}

// NewInterceptorLister returns a new InterceptorLister.
func NewInterceptorLister(indexer cache.Indexer) InterceptorLister {
	return &interceptorLister{indexer: indexer}
}

// List lists all Interceptors in the indexer.
func (s *interceptorLister) List(selector labels.Selector) (ret []*v1alpha1.Interceptor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Interceptor))
	})
	return ret, err
}

// Interceptors returns an object that can list and get Interceptors.
func (s *interceptorLister) Interceptors(namespace string) InterceptorNamespaceLister {
	return interceptorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InterceptorNamespaceLister helps list and get Interceptors.
// All objects returned here must be treated as read-only.
type InterceptorNamespaceLister interface {
	// List lists all Interceptors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Interceptor, err error)
	// Get retrieves the Interceptor from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Interceptor, error)
	InterceptorNamespaceListerExpansion
}

// interceptorNamespaceLister implements the InterceptorNamespaceLister
// interface.
type interceptorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Interceptors in the indexer for a given namespace.
func (s interceptorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Interceptor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Interceptor))
	})
	return ret, err
}

// Get retrieves the Interceptor from the indexer for a given namespace and name.
func (s interceptorNamespaceLister) Get(name string) (*v1alpha1.Interceptor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("interceptor"), name)
	}
	return obj.(*v1alpha1.Interceptor), nil
}
