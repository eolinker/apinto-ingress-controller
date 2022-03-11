package validation

import "github.com/eolinker/apinto-ingress-controller/pkg/apinto"

type admissionCheck interface {
	apinto.ProfessionCheck
}

type admissionChecker struct {
	cluster apinto.Cluster
}

var (
	validator admissionCheck
)

func InitAdmissionChecker(c *apinto.ClusterOptions) error {

	cluster, err := apinto.NewCluster(c)
	if err != nil {
		//TODO 打印日志

		return err
	}

	validator = &admissionChecker{
		cluster,
	}

	return nil
}

func (a *admissionChecker) RouterChecker() apinto.Checker {
	return a.RouterChecker()
}

func (a *admissionChecker) ServiceChecker() apinto.Checker {
	return a.ServiceChecker()
}

func (a *admissionChecker) UpstreamChecker() apinto.Checker {
	return a.UpstreamChecker()
}

func (a *admissionChecker) DiscoveryChecker() apinto.Checker {
	return a.DiscoveryChecker()
}

func (a *admissionChecker) OutputChecker() apinto.Checker {
	return a.OutputChecker()
}

func (a *admissionChecker) AuthChecker() apinto.Checker {
	return a.AuthChecker()
}

func (a *admissionChecker) SettingChecker() apinto.Checker {
	return a.SettingChecker()
}
