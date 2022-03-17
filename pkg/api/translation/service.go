package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeServiceToApinto(as *kubev1.ApintoService) *apintov1.Service {

	kService := as.Spec
	//拷贝Plugins
	plugins := make(map[string]apintov1.PluginConfig)
	for k, v := range kService.Plugins {
		plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
	}

	apintoService := &apintov1.Service{
		Metadata: apintov1.Metadata{
			Name:       kService.Name,
			Profession: "service",
			Driver:     kService.Driver,
			ID:         fmt.Sprintf("%s@service", kService.Name),
		},
		Timeout:     kService.Timeout,
		Retry:       kService.Retry,
		RewriteUrl:  kService.RewriteUrl,
		Scheme:      kService.Scheme,
		ProxyMethod: kService.ProxyMethod,
		Upstream:    kService.Upstream,
		Plugins:     plugins,
	}

	//若Anonymous不为空
	if kService.Anonymous != nil {
		apintoService.Anonymous = &apintov1.AnonymousConfig{Type: kService.Anonymous.Type, Config: kService.Anonymous.Config}
	}

	return apintoService
}
