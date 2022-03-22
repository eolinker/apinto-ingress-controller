package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeDiscoveryToApinto(ad *kubev1.ApintoDiscovery) *apintov1.Discovery {
	kDiscovery := ad.Spec

	apintoDiscovery := &apintov1.Discovery{
		Metadata: apintov1.Metadata{
			Name:       kDiscovery.Name,
			Profession: "discovery",
			Driver:     kDiscovery.Driver,
			ID:         fmt.Sprintf("%s@discovery", kDiscovery.Name),
		},
		Scheme:   kDiscovery.Scheme,
		HealthON: kDiscovery.HealthON,
		Config:   apintov1.Config(kDiscovery.Config),
		Health:   apintov1.HealthConfig(kDiscovery.Health),
	}

	return apintoDiscovery
}
