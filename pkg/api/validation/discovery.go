package validation

import (
	"context"
	"errors"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

var (
	errNotApintoDiscovery = errors.New("object is not ApintoDiscovery")
)

var ApintoDiscoveryValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成discovery
		ad, ok := object.(*kubev1.ApintoDiscovery)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, nil
		}
		discoveryList := ad.Spec

		switch review.Operation {
		case "delete":

		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
