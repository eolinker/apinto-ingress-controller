package validation

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession"
	"github.com/eolinker/eosc/log"
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

func InitAdmission(cfg *cluster.ClusterOptions) error {

	c, err := cluster.NewCluster(cfg)
	if err != nil {
		log.Errorf("failed to init admissionServer: %s", err)
		return err
	}

	validator = &admission{
		c,
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
