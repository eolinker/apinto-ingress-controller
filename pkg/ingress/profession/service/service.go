package service

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type serviceController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (s *serviceController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *serviceController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *serviceController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *serviceController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
