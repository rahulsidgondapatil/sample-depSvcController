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

package v1alpha1

import (
	v1alpha1 "github.com/infracloudio/sample-depSvcController/pkg/apis/SampleDepSvcController/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DepSvcResourceLister helps list DepSvcResources.
type DepSvcResourceLister interface {
	// List lists all DepSvcResources in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DepSvcResource, err error)
	// DepSvcResources returns an object that can list and get DepSvcResources.
	DepSvcResources(namespace string) DepSvcResourceNamespaceLister
	DepSvcResourceListerExpansion
}

// depSvcResourceLister implements the DepSvcResourceLister interface.
type depSvcResourceLister struct {
	indexer cache.Indexer
}

// NewDepSvcResourceLister returns a new DepSvcResourceLister.
func NewDepSvcResourceLister(indexer cache.Indexer) DepSvcResourceLister {
	return &depSvcResourceLister{indexer: indexer}
}

// List lists all DepSvcResources in the indexer.
func (s *depSvcResourceLister) List(selector labels.Selector) (ret []*v1alpha1.DepSvcResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DepSvcResource))
	})
	return ret, err
}

// DepSvcResources returns an object that can list and get DepSvcResources.
func (s *depSvcResourceLister) DepSvcResources(namespace string) DepSvcResourceNamespaceLister {
	return depSvcResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DepSvcResourceNamespaceLister helps list and get DepSvcResources.
type DepSvcResourceNamespaceLister interface {
	// List lists all DepSvcResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DepSvcResource, err error)
	// Get retrieves the DepSvcResource from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DepSvcResource, error)
	DepSvcResourceNamespaceListerExpansion
}

// depSvcResourceNamespaceLister implements the DepSvcResourceNamespaceLister
// interface.
type depSvcResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DepSvcResources in the indexer for a given namespace.
func (s depSvcResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DepSvcResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DepSvcResource))
	})
	return ret, err
}

// Get retrieves the DepSvcResource from the indexer for a given namespace and name.
func (s depSvcResourceNamespaceLister) Get(name string) (*v1alpha1.DepSvcResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("depsvcresource"), name)
	}
	return obj.(*v1alpha1.DepSvcResource), nil
}
