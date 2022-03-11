package validation

import (
	"context"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApintoGlobalSettingValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成globalSetting
		as, ok := object.(*kubev1.ApintoSetting)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		kSetting := as.Spec

		switch review.Operation {
		case "create", "update":
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

			apintoService := apintov1.Setting{
				Name:       "plugin",
				Profession: "setting",
				Driver:     "plugin",
				Plugins:    plugins,
			}

			_, err = validator.SettingChecker().UpdateCheck("plugin", apintoService)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
