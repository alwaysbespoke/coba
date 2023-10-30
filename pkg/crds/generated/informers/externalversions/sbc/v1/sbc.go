/*
Copyright 2023 Eric Hicks

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

	versioned "github.com/alwaysbespoke/coba/pkg/crds/generated/clientset/versioned"
	internalinterfaces "github.com/alwaysbespoke/coba/pkg/crds/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/alwaysbespoke/coba/pkg/crds/generated/listers/sbc/v1"
	sbcv1 "github.com/alwaysbespoke/coba/pkg/crds/sbc/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SBCInformer provides access to a shared informer and lister for
// SBCs.
type SBCInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SBCLister
}

type sBCInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSBCInformer constructs a new informer for SBC type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSBCInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSBCInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSBCInformer constructs a new informer for SBC type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSBCInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SbcV1().SBCs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SbcV1().SBCs(namespace).Watch(context.TODO(), options)
			},
		},
		&sbcv1.SBC{},
		resyncPeriod,
		indexers,
	)
}

func (f *sBCInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSBCInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sBCInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sbcv1.SBC{}, f.defaultInformer)
}

func (f *sBCInformer) Lister() v1.SBCLister {
	return v1.NewSBCLister(f.Informer().GetIndexer())
}
