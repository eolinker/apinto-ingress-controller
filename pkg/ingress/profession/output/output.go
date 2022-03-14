package output

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type outputController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (o *outputController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (o *outputController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (o *outputController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (o *outputController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (o *outputController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (o *outputController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
