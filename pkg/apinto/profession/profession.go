package profession

import (
	"context"
	"errors"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

var ErrClusterNotExist = errors.New("cluster not exist")

type ProfessionCheck interface {
	RouterChecker() Checker
	ServiceChecker() Checker
	UpstreamChecker() Checker
	DiscoveryChecker() Checker
	OutputChecker() Checker
	AuthChecker() Checker
	SettingChecker() Checker
}

type Profession interface {
	Router() Router
	Upstream() Upstream
	Service() Service
	Discovery() Discovery
	Output() Output
	Auth() Auth
	Setting() Setting
}

type Lister interface {
	List(ctx context.Context) ([]*response.Response, error)
}

type Checker interface {
	DelCheck(name string) (*response.Response, error)
	UpdateCheck(name string, value interface{}) (*response.Response, error)
}
type Service interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, service *v1.Service) (string, error)
	Create(ctx context.Context, service *v1.Service) (string, error)
}

type Router interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Router, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, router *v1.Router) (string, error)
	Create(ctx context.Context, router *v1.Router) (string, error)
}
type Upstream interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Upstream, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, upstream *v1.Upstream) (string, error)
	Create(ctx context.Context, upstream *v1.Upstream) (string, error)
}

type Output interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Output, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, output *v1.Output) (string, error)
	Create(ctx context.Context, output *v1.Output) (string, error)
}

type Auth interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Auth, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, auth *v1.Auth) (string, error)
	Create(ctx context.Context, auth *v1.Auth) (string, error)
}

type Discovery interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Discovery, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, discovery *v1.Discovery) (string, error)
	Create(ctx context.Context, discovery *v1.Discovery) (string, error)
}

// Setting 全局插件配置，只有更新和查询
type Setting interface {
	GetPlugin(ctx context.Context) (*v1.Setting, error)
	UpdatePlugin(ctx context.Context, setting *v1.Setting) (string, error)
}
