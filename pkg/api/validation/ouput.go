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
	"reflect"
	"strings"
)

var ApintoOutputValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msg string
		//将object转化成output
		ao, ok := object.(*kubev1.ApintoOutput)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		kOutput := ao.Spec

		switch review.Operation {
		case "create", "update":

			apintoOutput := apintov1.Output{
				Metadata: apintov1.Metadata{
					Name:       kOutput.Name,
					Profession: "output",
					Driver:     kOutput.Driver,
					ID:         fmt.Sprintf("%s@output", kOutput.Name),
				},
			}

			switch kOutput.Driver {
			case "file_output":
				fileConf := kOutput.FileOutput.Config
				cfg := make(apintov1.Config)
				cfg["file"] = fileConf.File
				cfg["dir"] = fileConf.Dir
				cfg["period"] = fileConf.Period
				cfg["expire"] = fileConf.Expire
				cfg["type"] = fileConf.Type
				cfg["formatter"] = map[string][]string(fileConf.Formatter)

				apintoOutput.Config = cfg
			case "nsqd":
				nsqConf := kOutput.Nsqd.Config
				cfg := make(apintov1.Config)
				cfg["topic"] = nsqConf.Topic
				cfg["address"] = nsqConf.Address
				cfg["nsq_conf"] = nsqConf.ClientConf
				cfg["type"] = nsqConf.Type
				cfg["formatter"] = map[string][]string(nsqConf.Formatter)

				apintoOutput.Config = cfg
			case "http_output":
				httpConf := kOutput.HttpOutput.Config
				cfg := make(apintov1.Config)
				cfg["method"] = httpConf.Method
				cfg["url"] = httpConf.Url
				cfg["headers"] = httpConf.Headers
				cfg["type"] = httpConf.Type
				cfg["formatter"] = map[string][]string(httpConf.Formatter)

				apintoOutput.Config = cfg
			case "syslog_output":
				sysConf := kOutput.SysOutput.Config
				cfg := make(apintov1.Config)
				cfg["network"] = sysConf.Network
				cfg["address"] = sysConf.Address
				cfg["level"] = sysConf.Level
				cfg["type"] = sysConf.Type
				cfg["formatter"] = map[string][]string(sysConf.Formatter)

				apintoOutput.Config = cfg
			case "kafka_output":
				kafkaConf := kOutput.KafkaOutput.Config
				cfg := make(apintov1.Config)
				cfg["topic"] = kafkaConf.Topic
				cfg["address"] = kafkaConf.Address
				cfg["timeout"] = kafkaConf.Timeout
				cfg["version"] = kafkaConf.Version
				cfg["partition_type"] = kafkaConf.PartitionType
				cfg["partition"] = kafkaConf.Partition
				cfg["partition_key"] = kafkaConf.PartitionKey
				cfg["type"] = kafkaConf.Type
				cfg["formatter"] = map[string][]string(kafkaConf.Formatter)

				apintoOutput.Config = cfg
			}

			_, err = validator.DiscoveryChecker().UpdateCheck(kOutput.Name, apintoOutput)
			if err != nil {
				valid = false
				msg = err.Error()
			}

		case "delete":
			_, err = validator.OutputChecker().DelCheck(kOutput.Name)
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
