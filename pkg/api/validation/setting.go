package validation

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/translation"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	errNotApintoSetting = errors.New("object is not ApintoSetting")
)

var ApintoGlobalSettingValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string

		switch review.Operation {
		case "create", "update":
			as := &kubev1.ApintoGlobalSetting{}
			err = json.Unmarshal(review.NewObjectRaw, as)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoSetting.Error()}, nil
			}

			apintoSetting := translation.KubeSettingToApinto(as)

			_, err = validator.Setting().UpdatePlugin(ctx, apintoSetting)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "delete":
			apintoSetting := &apintov1.Setting{
				Name:       "plugin",
				Profession: "setting",
				Driver:     "plugin",
				Plugins:    apintov1.SettingPlugins{},
			}

			_, err = validator.Setting().UpdatePlugin(ctx, apintoSetting)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)
