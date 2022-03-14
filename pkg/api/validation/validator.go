package validation

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/profession"
	"github.com/eolinker/eosc/log"
)

type admissionCheck interface {
	profession.ProfessionCheck
}

type admissionChecker struct {
	cluster cluster.Cluster
}

var (
	validator admissionCheck
)

func InitAdmissionChecker(cfg *cluster.ClusterOptions) error {

	c, err := cluster.NewCluster(cfg)
	if err != nil {
		log.Errorf("failed to init admissionChecker: %s", err)
		return err
	}

	validator = &admissionChecker{
		c,
	}

	return nil
}

func (a *admissionChecker) RouterChecker() profession.Checker {
	return a.RouterChecker()
}

func (a *admissionChecker) ServiceChecker() profession.Checker {
	return a.ServiceChecker()
}

func (a *admissionChecker) UpstreamChecker() profession.Checker {
	return a.UpstreamChecker()
}

func (a *admissionChecker) DiscoveryChecker() profession.Checker {
	return a.DiscoveryChecker()
}

func (a *admissionChecker) OutputChecker() profession.Checker {
	return a.OutputChecker()
}

func (a *admissionChecker) AuthChecker() profession.Checker {
	return a.AuthChecker()
}

func (a *admissionChecker) SettingChecker() profession.Checker {
	return a.SettingChecker()
}
