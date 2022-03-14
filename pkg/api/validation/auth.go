package validation

import (
	"context"
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApintoAuthValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成auth
		aa, ok := object.(*kubev1.ApintoAuth)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		kAuth := aa.Spec

		switch review.Operation {
		case "create", "update":

			apintoAuth := apintov1.Auth{
				Metadata: apintov1.Metadata{
					Name:       kAuth.Name,
					Profession: "auth",
					Driver:     kAuth.Driver,
					ID:         fmt.Sprintf("%s@auth", kAuth.Name),
				},
				HideCredentials:   kAuth.HideCredentials,
				RunOnPreflight:    kAuth.RunOnPreflight,
				SignatureIsBase64: kAuth.SignatureIsBase64,
				ClaimsToVerify:    kAuth.ClaimsToVerify,
			}

			if kAuth.User != nil {
				userSlice := make(apintov1.Users, 0, len(kAuth.User))
				for _, v := range kAuth.User {
					userSlice = append(userSlice, apintov1.User(v))
				}
				apintoAuth.User = userSlice
			}

			if kAuth.Credentials != nil {
				credentials := make(apintov1.Credentials, 0, len(kAuth.Credentials))
				for _, v := range kAuth.Credentials {
					credentials = append(credentials, apintov1.Credential(v))
				}
				apintoAuth.Credentials = credentials
			}

			_, err = validator.AuthChecker().UpdateCheck(kAuth.Name, apintoAuth)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			_, err = validator.AuthChecker().DelCheck(kAuth.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
