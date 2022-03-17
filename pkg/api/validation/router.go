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

		switch review.Operation {
		case "create":
			apintoRouter := translation.KubeRouterToApinto(ar)
			_, err = validator.Router().Create(ctx, apintoRouter)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "update":
			apintoRouter := translation.KubeRouterToApinto(ar)
			_, err = validator.Router().Update(ctx, apintoRouter)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "delete":
			err = validator.Router().Delete(ctx, ar.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}

		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, nil
	},
)
