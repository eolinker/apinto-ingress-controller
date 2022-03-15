package validation

import (
	"context"
	"errors"
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoDiscovery = errors.New("object is not ApintoDiscovery")
)

var ApintoDiscoveryValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		//将object转化成discovery
		ad, ok := object.(*kubev1.ApintoDiscovery)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, errNotApintoDiscovery
		}
		kDiscovery := ad.Spec

		switch review.Operation {
		case "create", "update":

			apintoDiscovery := apintov1.Discovery{
				Metadata: apintov1.Metadata{
					Name:       kDiscovery.Name,
					Profession: "discovery",
					Driver:     kDiscovery.Driver,
					ID:         fmt.Sprintf("%s@discovery", kDiscovery.Name),
				},
				Scheme:   kDiscovery.Scheme,
				HealthON: kDiscovery.HealthON,
				Config:   apintov1.Config(kDiscovery.Config),
				Health:   apintov1.HealthConfig(kDiscovery.Health),
			}

			_, err = validator.DiscoveryChecker().UpdateCheck(kDiscovery.Name, apintoDiscovery)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			_, err = validator.DiscoveryChecker().DelCheck(kDiscovery.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
