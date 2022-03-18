package validation

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/translation"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	"github.com/eolinker/eosc/log"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoAuth = errors.New("object is not ApintoAuth")
)

var ApintoAuthValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成auth
		//aa, ok := object.(*kubev1.ApintoAuth)
		//if !ok {
		//	return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoAuth.Error()}, nil
		//}
		log.Info(review)
		switch review.Operation {
		case "create":
			aa := &kubev1.ApintoAuth{}
			err = json.Unmarshal(review.NewObjectRaw, aa)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoAuth.Error()}, nil
			}

			apintoAuth := translation.KubeAuthToApinto(aa)

			_, err = validator.Auth().Create(ctx, apintoAuth)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "update":
			aa := &kubev1.ApintoAuth{}
			err = json.Unmarshal(review.NewObjectRaw, aa)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoAuth.Error()}, nil
			}

			apintoAuth := translation.KubeAuthToApinto(aa)

			_, err = validator.Auth().Update(ctx, apintoAuth)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			aa := &kubev1.ApintoAuth{}
			err = json.Unmarshal(review.OldObjectRaw, aa)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoAuth.Error()}, nil
			}

			err = validator.Auth().Delete(ctx, aa.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
