package cluster

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/auth"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/discovery"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/output"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/router"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/service"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/setting"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession/upstream"
	"time"
)

// Cluster 集群服务接口
type Cluster interface {
	profession.ProfessionCheck
	profession.Profession
}
type ClusterOptions struct {
	Name     string
	AdminKey string
	BaseURL  string
	Timeout  time.Duration
}

type cluster struct {
	name      string
	client    client.Client
	router    *router.Router
	upstream  *upstream.Upstream
	service   *service.Service
	discovery *discovery.Discovery
	output    *output.Output
	auth      *auth.Auth
	setting   *setting.Setting
}

func (c *cluster) RouterChecker() profession.Checker {
	return c.router
}

func (c *cluster) ServiceChecker() profession.Checker {
	return c.service
}

func (c *cluster) UpstreamChecker() profession.Checker {
	return c.upstream
}

func (c *cluster) DiscoveryChecker() profession.Checker {
	return c.upstream
}

func (c *cluster) OutputChecker() profession.Checker {
	return c.output
}

func (c *cluster) AuthChecker() profession.Checker {
	return c.auth
}

func (c *cluster) SettingChecker() profession.Checker {
	return c.setting
}

func NewCluster(c *ClusterOptions) (*cluster, error) {
	cli, err := client.NewClient(c.BaseURL, c.Timeout, c.AdminKey)
	if err != nil {
		return nil, err
	}
	return &cluster{
		name:      c.Name,
		client:    cli,
		router:    router.NewRouter(cli),
		service:   service.NewService(cli),
		upstream:  upstream.NewUpstream(cli),
		auth:      auth.NewAuth(cli),
		output:    output.NewOutput(cli),
		discovery: discovery.NewDiscovery(cli),
		setting:   setting.NewSetting(cli),
	}, nil
}

func (c *cluster) Router() profession.Router {
	return c.router
}

func (c *cluster) Upstream() profession.Upstream {
	return c.upstream
}

func (c *cluster) Service() profession.Service {
	return c.service
}

func (c *cluster) Discovery() profession.Discovery {
	return c.discovery
}

func (c *cluster) Output() profession.Output {
	return c.output
}

func (c *cluster) Auth() profession.Auth {
	return c.auth
}

func (c *cluster) Setting() profession.Setting {
	return c.setting
}
