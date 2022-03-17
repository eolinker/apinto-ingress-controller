package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeRouterToApinto(ar *kubev1.ApintoRouter) *apintov1.Router {
	kRouter := ar.Spec

	//拷贝Cert
	cert := make([]apintov1.Cert, 0, len(kRouter.Cert))
	for i, v := range kRouter.Cert {
		cert[i] = apintov1.Cert{Key: v.Key, Crt: v.Crt}
	}

	//拷贝rules
	rules := make([]apintov1.Rule, 0, len(kRouter.Rules))
	for i, v := range kRouter.Rules {
		rules[i] = apintov1.Rule{Location: v.Location, Header: v.Header, Query: v.Query}
	}

	//拷贝Plugins
	plugins := make(map[string]apintov1.PluginConfig)
	for k, v := range kRouter.Plugins {
		plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
	}

	apintoRouter := &apintov1.Router{
		Metadata: apintov1.Metadata{
			Name:       kRouter.Name,
			Profession: "router",
			Driver:     kRouter.Driver,
			ID:         fmt.Sprintf("%s@router", kRouter.Name),
		},
		Listen:   kRouter.Listen,
		Target:   kRouter.Target,
		Method:   kRouter.Method,
		Host:     kRouter.Host,
		Protocol: kRouter.Protocol,
		Cert:     cert,
		Rules:    rules,
		Plugins:  plugins,
	}

	return apintoRouter
}
