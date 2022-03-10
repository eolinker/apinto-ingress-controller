package apinto

import (
	"context"
	"errors"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
	"io"
	"sync"
)

var (
	_errClusterNotExist   = errors.New("cluster not exist")
	_errClusterDuplicated = errors.New("register cluster duplicated")
)

type Apinto interface {
	Cluster(name string) (Cluster, error)
	AddCluster(*ClusterOptions) error
	UpdateCluster(*ClusterOptions) error
	ListClusters() []Cluster
	DeleteCluster(name string) error
}

// Cluster 集群服务接口
type Cluster interface {
	Router() Router
	Upstream() Upstream
	Service() Service
	Discovery() Discovery
	Output() Output
	Auth() Auth
	Setting() Setting
}

// Client 发送apinto接口请求
// 发送请求时注意是否有admin key
type Client interface {
	Url() string
	Get(ctx context.Context, url string) ([]byte, error)
	List(ctx context.Context, url string) ([]*response.Response, error)
	Create(ctx context.Context, url string, body io.Reader) (*response.Response, error)
	Delete(ctx context.Context, url string) (*response.Response, error)
	Update(ctx context.Context, url string, body io.Reader) (*response.Response, error)
}
type Lister interface {
	List(ctx context.Context) ([]*response.Response, error)
}
type Router interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Router, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, router *v1.Router) (string, error)
	Create(ctx context.Context, router *v1.Router) (string, error)
}
type Service interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Service, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, service *v1.Service) (string, error)
	Create(ctx context.Context, service *v1.Service) (string, error)
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

type Upstream interface {
	Lister
	Get(ctx context.Context, name string) (*v1.Upstream, error)
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, upstream *v1.Upstream) (string, error)
	Create(ctx context.Context, upstream *v1.Upstream) (string, error)
}

type apinto struct {
	mu   sync.RWMutex
	data map[string]Cluster
}

func NewApinto() (Apinto, error) {
	cli := &apinto{
		data: make(map[string]Cluster),
	}
	return cli, nil
}

func (a *apinto) Cluster(name string) (Cluster, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	c, ok := a.data[name]
	if !ok {
		return nil, _errClusterNotExist
	}
	return c, nil
}

func (a *apinto) AddCluster(options *ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	_, ok := a.data[options.Name]
	if ok {
		return _errClusterDuplicated
	}
	c, err := NewCluster(options)
	if err != nil {
		return err
	}
	a.data[options.Name] = c
	return nil
}

func (a *apinto) UpdateCluster(options *ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.data[options.Name]; !ok {
		return _errClusterNotExist
	}
	c, err := NewCluster(options)
	if err != nil {
		return err
	}

	a.data[options.Name] = c
	return nil
}

func (a *apinto) ListClusters() []Cluster {
	a.mu.RLock()
	defer a.mu.RUnlock()
	clusters := make([]Cluster, 0, len(a.data))
	for _, c := range a.data {
		clusters = append(clusters, c)
	}
	return clusters
}

func (a *apinto) DeleteCluster(name string) error {
	a.mu.Lock()
	delete(a.data, name)
	a.mu.Unlock()
	return nil
}
