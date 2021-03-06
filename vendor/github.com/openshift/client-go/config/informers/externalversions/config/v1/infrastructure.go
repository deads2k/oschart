// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	config_v1 "github.com/openshift/api/config/v1"
	versioned "github.com/openshift/client-go/config/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/config/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/config/listers/config/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InfrastructureInformer provides access to a shared informer and lister for
// Infrastructures.
type InfrastructureInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.InfrastructureLister
}

type infrastructureInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewInfrastructureInformer constructs a new informer for Infrastructure type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInfrastructureInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInfrastructureInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredInfrastructureInformer constructs a new informer for Infrastructure type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInfrastructureInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().Infrastructures().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1().Infrastructures().Watch(options)
			},
		},
		&config_v1.Infrastructure{},
		resyncPeriod,
		indexers,
	)
}

func (f *infrastructureInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInfrastructureInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *infrastructureInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&config_v1.Infrastructure{}, f.defaultInformer)
}

func (f *infrastructureInformer) Lister() v1.InfrastructureLister {
	return v1.NewInfrastructureLister(f.Informer().GetIndexer())
}
