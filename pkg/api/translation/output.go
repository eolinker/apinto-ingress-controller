package translation

import (
	"fmt"
	kubev1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/configs/apinto/v1"
	apintov1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

func KubeOutputToApinto(ao *kubev1.ApintoOutput) *apintov1.Output {
	kOutput := ao.Spec

	apintoOutput := &apintov1.Output{
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

	return apintoOutput
}
