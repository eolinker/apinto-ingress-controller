package validation

import (
	"context"
	"errors"
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoRouter = errors.New("object is not ApintoRouter")
)

var ApintoRouterValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {

		valid := true
		var msg string

		ar, ok := object.(*kubev1.ApintoRouter)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoRouter.Error()}, errNotApintoRouter
		}
		kRouter := ar.Spec

		switch review.Operation {
		case "create", "update":

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

			apintoRouter := apintov1.Router{
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

			_, err := validator.RouterChecker().UpdateCheck(kRouter.Name, apintoRouter)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}

		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, nil
	},
)
