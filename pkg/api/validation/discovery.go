package validation

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/translation"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoDiscovery = errors.New("object is not ApintoDiscovery")
)

var ApintoDiscoveryValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		switch review.Operation {
		case "create":
			ad := &kubev1.ApintoDiscovery{}
			err = json.Unmarshal(review.NewObjectRaw, ad)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, nil
			}

			apintoDiscovery := translation.KubeDiscoveryToApinto(ad)

			_, err = validator.Discovery().Create(ctx, apintoDiscovery)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "update":
			ad := &kubev1.ApintoDiscovery{}
			err = json.Unmarshal(review.NewObjectRaw, ad)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, nil
			}

			apintoDiscovery := translation.KubeDiscoveryToApinto(ad)

			_, err = validator.Discovery().Update(ctx, apintoDiscovery)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			ad := &kubev1.ApintoDiscovery{}
			err = json.Unmarshal(review.OldObjectRaw, ad)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, nil
			}

			err = validator.Discovery().Delete(ctx, ad.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
