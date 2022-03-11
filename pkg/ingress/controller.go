package ingress

import (
	"context"

	"github.com/eolinker/apinto-ingress-controller/pkg/types"
)

type IIngressControllerFactory interface {
	// NewController create CRD controller,name is same as kubernetes CRD kind
	// workers means the number of goroutines that handle tasks concurrently
	NewController(name string, workers int) IIngressController
}

type IIngressController interface {
	// Run controller will run and watch kubernetes api server changes.
	Run(ctx context.Context) error
	// Sync  changes to apinto cluster
	Sync(ctx context.Context, ev *types.Event) error
	// HandleSyncErr handle sync error.while the error occur,obj will push workqueue again
	HandleSyncErr(obj interface{}, errOrigin error)
	IIngressControllerEvent
}
