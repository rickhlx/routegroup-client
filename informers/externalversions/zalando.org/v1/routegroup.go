//lint:file-ignore ST1005 Generated code

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	zalandoorgv1 "github.com/szuecs/routegroup-client/apis/zalando.org/v1"
	internalinterfaces "github.com/szuecs/routegroup-client/informers/externalversions/internalinterfaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	v1 "k8s.io/kubernetes/pkg/client/listers/zalando.org/v1"
)

// RouteGroupInformer provides access to a shared informer and lister for
// RouteGroups.
type RouteGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.RouteGroupLister
}

type routeGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRouteGroupInformer constructs a new informer for RouteGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRouteGroupInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRouteGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRouteGroupInformer constructs a new informer for RouteGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRouteGroupInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ZalandoV1().RouteGroups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ZalandoV1().RouteGroups(namespace).Watch(context.TODO(), options)
			},
		},
		&zalandoorgv1.RouteGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *routeGroupInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRouteGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *routeGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&zalandoorgv1.RouteGroup{}, f.defaultInformer)
}

func (f *routeGroupInformer) Lister() v1.RouteGroupLister {
	return v1.NewRouteGroupLister(f.Informer().GetIndexer())
}
