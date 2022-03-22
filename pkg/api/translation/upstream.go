package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeUpstreamToApinto(au *kubev1.ApintoUpstream) *apintov1.Upstream {
	kUpstream := au.Spec

	//拷贝Plugins
	plugins := make(map[string]apintov1.PluginConfig)
	for k, v := range kUpstream.Plugins {
		plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
	}

	apintoUpstream := &apintov1.Upstream{
		Metadata: apintov1.Metadata{
			Name:       kUpstream.Name,
			Profession: "upstream",
			Driver:     kUpstream.Driver,
			ID:         fmt.Sprintf("%s@upstream", kUpstream.Name),
		},
		Discovery: kUpstream.Discovery,
		Config:    kUpstream.Config,
		Scheme:    kUpstream.Scheme,
		Type:      kUpstream.Type,
		Plugins:   plugins,
	}

	return apintoUpstream
}
