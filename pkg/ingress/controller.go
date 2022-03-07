package ingress

import (
	"context"

	"github.com/eolinker/apinto-ingress-controller/pkg/types"
)

type IIngressControllerFactory interface {
	// create CRD controller,name is same as kubernetes CRD kind
	// workers means the number of goroutines that handle tasks concurrently
	newController(name string, workers int) IIngressController
}

type IIngressController interface {
	// controller will run and watch kubernetes api server changes.
	run(ctx context.Context) error
	// sync changes to apinto cluster
	sync(ctx context.Context, ev *types.Event) error
	// handle sync error.while the error occur,obj will push workqueue again
	handleSyncErr(obj interface{}, errOrigin error)
	IIngressControllerEvent
}
