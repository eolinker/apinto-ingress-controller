package validation

import (
	"context"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var globalSettingListUrl = ""

func SetGlobalSettingListUrl(baseURl string) {
	globalSettingListUrl = baseURl + "/api/setting/plugin"
}

var ApintoGlobalSettingValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true

		switch review.Operation {
		case "create":
			// TODO 创建前检查是否已经存在有一个对象
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: ""}, nil
	},
)
