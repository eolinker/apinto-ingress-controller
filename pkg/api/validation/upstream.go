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
	errNotApintoUpstream = errors.New("object is not ApintoUpstream")
)

var ApintoUpstreamValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		//将object转化成upstream
		au, ok := object.(*kubev1.ApintoUpstream)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, errNotApintoUpstream
		}
		kUpstream := au.Spec

		switch review.Operation {
		case "create", "update":

			//拷贝Plugins
			plugins := make(map[string]apintov1.PluginConfig)
			for k, v := range kUpstream.Plugins {
				plugins[k] = apintov1.PluginConfig{Disable: v.Disable, Config: v.Config}
			}

			apintoUpstream := apintov1.Upstream{
				Metadata: apintov1.Metadata{
					Name:       kUpstream.Name,
					Profession: "upstream",
					Driver:     kUpstream.Driver,
				},
				Discovery: kUpstream.Discovery,
				Config:    kUpstream.Config,
				Scheme:    kUpstream.Scheme,
				Type:      kUpstream.Type,
				Plugins:   plugins,
			}

			_, err = validator.UpstreamChecker().UpdateCheck(kUpstream.Name, apintoUpstream)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			_, err = validator.UpstreamChecker().DelCheck(kUpstream.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
