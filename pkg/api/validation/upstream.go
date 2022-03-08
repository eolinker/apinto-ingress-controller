package validation

import (
	"context"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var upstreamListUrl = ""

func SetUpstreamListUrl(baseURl string) {
	upstreamListUrl = baseURl + "/api/upstream"
}

var ApintoUpstreamValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		return nil, nil
	},
)
