package validation

import (
	"context"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var discoveryListUrl = ""

func SetDiscoveryListUrl(baseURl string) {
	discoveryListUrl = baseURl + "/api/discovery"
}

var ApintoDiscoveryValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true

		switch review.Operation {
		case "delete":
			//TODO 检查是否有upstream在使用这个discovery
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: ""}, nil
	},
)
