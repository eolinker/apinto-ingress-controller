package validation

import (
	"context"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApintoGlobalSettingValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true

		//将object转化成globalSetting
		as, ok := object.(*kubev1.ApintoSetting)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		setting := as.Spec

		switch review.Operation {
		case "create":
			// TODO 创建前检查是否已经存在有一个对象
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: ""}, nil
	},
)
