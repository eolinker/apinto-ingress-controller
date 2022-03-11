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
	errNotApintoRouter = errors.New("object is not ApintoRouter")
)

var ApintoRouterValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {

		valid := true
		var msgs []string

		ar, ok := object.(*kubev1.ApintoRouter)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoRouter.Error()}, errNotApintoRouter
		}
		router := ar.Spec

		switch review.Operation {
		case "create", "update":

			//拷贝Cert
			cert := make([]apintov1.Cert, 0, len(router.Cert))
			for i, v := range router.Cert {
				cert[i] = apintov1.Cert{Key: v.Key, Crt: v.Crt}
			}

			//拷贝rules
			rules := make([]apintov1.Rule, 0, len(router.Rules))
			for i, v := range router.Rules {
				rules[i] = apintov1.Rule{Location: v.Location, Header: v.Header, Query: v.Query}
			}

			//拷贝Plugins
			plugins := make(map[string]apintov1.PluginConfig)
			for k, v := range router.Plugins {
				plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
			}

			apintoRouter := apintov1.Router{
				Metadata: apintov1.Metadata{
					Name:       router.Name,
					Profession: "router",
					Driver:     router.Driver,
				},
				Listen:   router.Listen,
				Target:   router.Target,
				Method:   router.Method,
				Host:     router.Host,
				Protocol: router.Protocol,
				Cert:     cert,
				Rules:    rules,
				Plugins:  plugins,
			}

			_, err := validator.RouterChecker().UpdateCheck(router.Name, apintoRouter)
			if err != nil {
				valid = false
			}
		}

		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
