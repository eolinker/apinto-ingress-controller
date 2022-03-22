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
	"context"
	time "time"

	versioned "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/clientset/versioned"
	internalinterfaces "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/informers/externalversions/internalinterfaces"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/listers/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApintoDiscoveryInformer provides access to a shared informer and lister for
// ApintoDiscoveries.
type ApintoDiscoveryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ApintoDiscoveryLister
}

type apintoDiscoveryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApintoDiscoveryInformer constructs a new informer for ApintoDiscovery type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApintoDiscoveryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApintoDiscoveryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApintoDiscoveryInformer constructs a new informer for ApintoDiscovery type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApintoDiscoveryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApintoV1().ApintoDiscoveries(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApintoV1().ApintoDiscoveries(namespace).Watch(context.TODO(), options)
			},
		},
		&apintov1.ApintoDiscovery{},
		resyncPeriod,
		indexers,
	)
}

func (f *apintoDiscoveryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApintoDiscoveryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *apintoDiscoveryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apintov1.ApintoDiscovery{}, f.defaultInformer)
}

func (f *apintoDiscoveryInformer) Lister() v1.ApintoDiscoveryLister {
	return v1.NewApintoDiscoveryLister(f.Informer().GetIndexer())
}
