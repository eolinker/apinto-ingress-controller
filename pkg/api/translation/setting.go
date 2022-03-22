package translation

import (
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeSettingToApinto(as *kubev1.ApintoGlobalSetting) *apintov1.Setting {
	kSetting := as.Spec

	//拷贝Plugins
	plugins := make(apintov1.SettingPlugins, 0, len(kSetting.Plugins))
	for _, v := range kSetting.Plugins {
		plugins = append(plugins, apintov1.SettingPlugin{
			ID:         v.ID,
			Name:       v.Name,
			Type:       v.Type,
			Status:     v.Status,
			Config:     apintov1.Config(v.Config),
			InitConfig: apintov1.Config(v.InitConfig),
		})
	}

	apintoSetting := &apintov1.Setting{
		Name:       "plugin",
		Profession: "setting",
		Driver:     "plugin",
		Plugins:    plugins,
	}

	return apintoSetting
}
