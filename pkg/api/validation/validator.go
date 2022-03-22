package validation

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
)

type admissionServer interface {
	profession.Profession
}

type admission struct {
	cluster cluster.Cluster
}

var (
	validator admissionServer
)

func InitAdmission(cfg *config.Config, apinto apinto.Apinto) error {
	validator = &admission{
		apinto.Cluster(cfg.APINTO.DefaultClusterName),
	}
	return nil
}

func (a *admission) Router() profession.Router {
	return a.cluster.Router()
}

func (a *admission) Service() profession.Service {
	return a.cluster.Service()
}

func (a *admission) Upstream() profession.Upstream {
	return a.cluster.Upstream()
}

func (a *admission) Discovery() profession.Discovery {
	return a.cluster.Discovery()
}

func (a *admission) Output() profession.Output {
	return a.cluster.Output()
}

func (a *admission) Auth() profession.Auth {
	return a.cluster.Auth()
}

func (a *admission) Setting() profession.Setting {
	return a.cluster.Setting()
}
