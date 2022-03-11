package discovery

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type discoveryController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (d *discoveryController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (d *discoveryController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (d *discoveryController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (d *discoveryController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d *discoveryController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (d *discoveryController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
