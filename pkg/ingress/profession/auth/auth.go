package auth

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type authController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (a *authController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a *authController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (a *authController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (a *authController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (a *authController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (a *authController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
