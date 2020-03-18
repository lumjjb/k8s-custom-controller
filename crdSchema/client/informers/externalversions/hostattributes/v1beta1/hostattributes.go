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

package v1beta1

import (
	time "time"

	hostattributesv1beta1 "github.com/intel-secl/k8s-custom-controller/crdSchema/api/hostattributes/v1beta1"
	versioned "github.com/intel-secl/k8s-custom-controller/crdSchema/client/clientset/versioned"
	internalinterfaces "github.com/intel-secl/k8s-custom-controller/crdSchema/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/intel-secl/k8s-custom-controller/crdSchema/client/listers/hostattributes/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HostAttributesInformer provides access to a shared informer and lister for
// HostAttributeses.
type HostAttributesInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.HostAttributesLister
}

type hostAttributesInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHostAttributesInformer constructs a new informer for HostAttributes type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHostAttributesInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHostAttributesInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHostAttributesInformer constructs a new informer for HostAttributes type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHostAttributesInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1beta1().HostAttributeses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CrdV1beta1().HostAttributeses(namespace).Watch(options)
			},
		},
		&hostattributesv1beta1.HostAttributes{},
		resyncPeriod,
		indexers,
	)
}

func (f *hostAttributesInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHostAttributesInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hostAttributesInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hostattributesv1beta1.HostAttributes{}, f.defaultInformer)
}

func (f *hostAttributesInformer) Lister() v1beta1.HostAttributesLister {
	return v1beta1.NewHostAttributesLister(f.Informer().GetIndexer())
}