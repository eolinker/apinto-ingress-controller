package validation

import (
	"context"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApintoOutputValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成output
		ao, ok := object.(*kubev1.ApintoOutput)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		kOutput := ao.Spec

		switch review.Operation {
		case "create", "update":
			//TODO

		case "delete":
			_, err = validator.OutputChecker().DelCheck(kOutput.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
