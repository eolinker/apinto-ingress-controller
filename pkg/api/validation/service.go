package validation

import (
	"context"
	"errors"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

var (
	errNotApintoService = errors.New("object is not ApintoService")
)

var ApintoServiceValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成service
		as, ok := object.(*kubev1.ApintoService)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoService.Error()}, errNotApintoService
		}
		service := as.Spec

		switch review.Operation {
		case "create", "update":
			//拷贝Anonymous
			anonymous := apintov1.AnonymousConfig{}

			//拷贝Plugins
			plugins := make(map[string]apintov1.PluginConfig)
			for k, v := range service.Plugins {
				plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
			}

			apintoService := apintov1.Service{
				Metadata: apintov1.Metadata{
					Name:       service.Name,
					Profession: "service",
					Driver:     service.Driver,
				},
				Timeout:     service.Timeout,
				Retry:       service.Retry,
				RewriteUrl:  service.RewriteUrl,
				Scheme:      service.Scheme,
				ProxyMethod: service.ProxyMethod,
				Upstream:    service.Upstream,

				Plugins: plugins,
			}

			_, err := validator.RouterChecker().UpdateCheck(service.Name, apintoService)
			if err != nil {
				valid = false
			}

		case "delete":

		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
