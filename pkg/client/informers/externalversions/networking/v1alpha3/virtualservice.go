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

package v1alpha3

import (
	time "time"

	networkingv1alpha3 "github.com/RuiWang14/k8s-istio-client/pkg/apis/networking/v1alpha3"
	versioned "github.com/RuiWang14/k8s-istio-client/pkg/client/clientset/versioned"
	internalinterfaces "github.com/RuiWang14/k8s-istio-client/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha3 "github.com/RuiWang14/k8s-istio-client/pkg/client/listers/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VirtualServiceInformer provides access to a shared informer and lister for
// VirtualServices.
type VirtualServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha3.VirtualServiceLister
}

type virtualServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualServiceInformer constructs a new informer for VirtualService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualServiceInformer constructs a new informer for VirtualService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().VirtualServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1alpha3().VirtualServices(namespace).Watch(options)
			},
		},
		&networkingv1alpha3.VirtualService{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1alpha3.VirtualService{}, f.defaultInformer)
}

func (f *virtualServiceInformer) Lister() v1alpha3.VirtualServiceLister {
	return v1alpha3.NewVirtualServiceLister(f.Informer().GetIndexer())
}
