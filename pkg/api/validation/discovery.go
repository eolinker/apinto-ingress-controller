package validation

import (
	"context"
	"errors"
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

		//将object转化成discovery
		ad, ok := object.(*kubev1.ApintoDiscovery)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, errNotApintoDiscovery
		}
		kDiscovery := ad.Spec

		switch review.Operation {
		case "delete":
			_, err = validator.DiscoveryChecker().DelCheck(kDiscovery.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
