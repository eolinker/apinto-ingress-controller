package validation

import (
	"context"
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/translation"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoUpstream = errors.New("object is not ApintoUpstream")
)

var ApintoUpstreamValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		//将object转化成upstream
		au, ok := object.(*kubev1.ApintoUpstream)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, errNotApintoUpstream
		}

		switch review.Operation {
		case "create":
			apintoUpstream := translation.KubeUpstreamToApinto(au)

			_, err = validator.Upstream().Create(ctx, apintoUpstream)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "update":
			apintoUpstream := translation.KubeUpstreamToApinto(au)

			_, err = validator.Upstream().Update(ctx, apintoUpstream)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			err = validator.Upstream().Delete(ctx, au.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
