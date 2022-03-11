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
	errNotApintoUpstream = errors.New("object is not ApintoUpstream")
)

var ApintoUpstreamValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成upstream
		au, ok := object.(*kubev1.ApintoUpstream)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		upstreamList := au.Spec

		switch review.Operation {
		case "create", "update":

		case "delete":

		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
