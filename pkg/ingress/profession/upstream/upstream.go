package upstream

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type upstreamController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (u *upstreamController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (u *upstreamController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (u *upstreamController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (u *upstreamController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u *upstreamController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (u *upstreamController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
