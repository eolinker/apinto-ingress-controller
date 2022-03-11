package validation

import (
	"context"
	"errors"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoService = errors.New("object is not ApintoService")
)

var ApintoServiceValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		//将object转化成service
		as, ok := object.(*kubev1.ApintoService)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoService.Error()}, errNotApintoService
		}
		kService := as.Spec

		switch review.Operation {
		case "create", "update":

			//拷贝Plugins
			plugins := make(map[string]apintov1.PluginConfig)
			for k, v := range kService.Plugins {
				plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
			}

			apintoService := apintov1.Service{
				Metadata: apintov1.Metadata{
					Name:       kService.Name,
					Profession: "service",
					Driver:     kService.Driver,
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

			_, err = validator.ServiceChecker().UpdateCheck(kService.Name, apintoService)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			_, err = validator.ServiceChecker().DelCheck(kService.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
