package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeAuthToApinto(aa *kubev1.ApintoAuth) *apintov1.Auth {
	kAuth := aa.Spec

	apintoAuth := &apintov1.Auth{
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

	return apintoAuth
}
