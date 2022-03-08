package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
)

// Setting 全局插件配置，只有更新和查询
type Setting interface {
	GetPlugin(ctx context.Context, name string) (*v1.Setting, error)
	UpdatePlugin(ctx context.Context, upstream *v1.Setting) (string, error)
}

type setting struct {
	client Client
	url    string
}

func NewSetting(client Client) *setting {
	return &setting{
		url:    fmt.Sprintf("%s/%s", client.Url(), "setting/plugin"),
		client: client,
	}
}

func (s *setting) GetPlugin(ctx context.Context, name string) (*v1.Setting, error) {
	//TODO implement me
	panic("implement me")
}

func (s *setting) UpdatePlugin(ctx context.Context, upstream *v1.Setting) (string, error) {
	//TODO implement me
	panic("implement me")
}
