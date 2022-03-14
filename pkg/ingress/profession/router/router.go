package router

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type routerController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (r *routerController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (r *routerController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (r *routerController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (r *routerController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (r *routerController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (r *routerController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
