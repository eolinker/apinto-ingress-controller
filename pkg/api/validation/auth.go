package validation

import (
	"context"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var authListUrl = ""

func SetAuthListUrl(baseURl string) {
	authListUrl = baseURl + "/api/auth"
}

var ApintoAuthValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true

		switch review.Operation {
		case "delete":
			//TODO 暂时不处理
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: ""}, nil
	},
)
