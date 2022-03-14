package setting

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"k8s.io/client-go/util/workqueue"
)

type settingController struct {
	name      string
	workqueue workqueue.RateLimitingInterface
	workers   int
}

func (s *settingController) Run(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *settingController) Sync(ctx context.Context, ev *types.Event) error {
	//TODO implement me
	panic("implement me")
}

func (s *settingController) HandleSyncErr(obj interface{}, errOrigin error) {
	//TODO implement me
	panic("implement me")
}

func (s *settingController) OnAdd(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *settingController) OnUpdate(obj interface{}) {
	//TODO implement me
	panic("implement me")
}

func (s *settingController) OnDelete(obj interface{}) {
	//TODO implement me
	panic("implement me")
}
