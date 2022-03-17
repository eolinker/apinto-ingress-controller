package validation

import (
	"context"
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/translation"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
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

		switch review.Operation {
		case "create":
			apintoService := translation.KubeServiceToApinto(as)
			_, err = validator.Service().Create(ctx, apintoService)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "update":
			apintoService := translation.KubeServiceToApinto(as)
			_, err = validator.Service().Update(ctx, apintoService)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			err = validator.Service().Delete(ctx, as.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
