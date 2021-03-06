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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	DepSvcResourcev1 "github.com/rahulsidgondapatil/sample-depSvcController/pkg/apis/DepSvcResource/v1"
	versioned "github.com/rahulsidgondapatil/sample-depSvcController/pkg/client/clientset/versioned"
	internalinterfaces "github.com/rahulsidgondapatil/sample-depSvcController/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/rahulsidgondapatil/sample-depSvcController/pkg/client/listers/DepSvcResource/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DepSvcResourceInformer provides access to a shared informer and lister for
// DepSvcResources.
type DepSvcResourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DepSvcResourceLister
}

type depSvcResourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDepSvcResourceInformer constructs a new informer for DepSvcResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDepSvcResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDepSvcResourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDepSvcResourceInformer constructs a new informer for DepSvcResource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDepSvcResourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DepSvcResourceV1().DepSvcResources(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DepSvcResourceV1().DepSvcResources(namespace).Watch(options)
			},
		},
		&DepSvcResourcev1.DepSvcResource{},
		resyncPeriod,
		indexers,
	)
}

func (f *depSvcResourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDepSvcResourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *depSvcResourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&DepSvcResourcev1.DepSvcResource{}, f.defaultInformer)
}

func (f *depSvcResourceInformer) Lister() v1.DepSvcResourceLister {
	return v1.NewDepSvcResourceLister(f.Informer().GetIndexer())
}
