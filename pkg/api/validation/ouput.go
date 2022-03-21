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
	"reflect"
	"strings"
)

var (
	errNotApintoOutput = errors.New("object is not ApintoOutput")
)

var ApintoOutputValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成output
		//ao, ok := object.(*kubev1.ApintoOutput)
		//if !ok {
		//	return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoOutput.Error()}, nil
		//}
		//log.Info(review)

		switch review.Operation {
		case "create":
			ao := &kubev1.ApintoOutput{}
			err = json.Unmarshal(review.NewObjectRaw, ao)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoOutput.Error()}, nil
			}

			apintoOutput := translation.KubeOutputToApinto(ao)
			_, err = validator.Output().Create(ctx, apintoOutput)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		case "update":
			ao := &kubev1.ApintoOutput{}
			err = json.Unmarshal(review.NewObjectRaw, ao)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoOutput.Error()}, nil
			}

			apintoOutput := translation.KubeOutputToApinto(ao)

			_, err = validator.Output().Update(ctx, apintoOutput)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			ao := &kubev1.ApintoOutput{}
			err = json.Unmarshal(review.OldObjectRaw, ao)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoOutput.Error()}, nil
			}

			err = validator.Output().Delete(ctx, ao.Spec.Name)
			if err != nil {
				valid = false
				msg = err.Error()
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: msg}, err
	},
)

func transToConfig(cfg interface{}) (apintov1.Config, error) {
	configs := make(apintov1.Config)

	ty := reflect.TypeOf(cfg)
	ptr := reflect.ValueOf(cfg)
	va := ptr.Elem()
	if ty.Kind() != reflect.Struct {
		return nil, errors.New("output config is not struct")
	}

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		fieldTag := field.Tag.Get("yaml")
		fieldTag = strings.Split(fieldTag, ",")[0]

		configs[fieldTag] = va.Field(i).Interface()

	}
	return configs, nil
}
