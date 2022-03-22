package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoOutput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OutputSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type OutputSpec struct {
	Name        string      `json:"name,omitempty" yaml:"name,omitempty"`
	Driver      string      `json:"driver,omitempty" yaml:"driver,omitempty"`
	FileOutput  FileOutput  `json:"file_output,omitempty" yaml:"file_output,omitempty"`
	Nsqd        NsqdOutput  `json:"nsqd,omitempty" yaml:"nsqd,omitempty"`
	HttpOutput  HttpOutput  `json:"http_output,omitempty" yaml:"http_output,omitempty"`
	SysOutput   SysOutput   `json:"syslog_output,omitempty" yaml:"syslog_output,omitempty"`
	KafkaOutput KafkaOutput `json:"kafka_output,omitempty" yaml:"kafka_output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoOutputList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoOutput `json:"items,omitempty" yaml:"items,omitempty"`
}

type FileOutput struct {
	Config FileConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
type FileConfig struct {
	File      string          `json:"file,omitempty" yaml:"file,omitempty"`
	Dir       string          `json:"dir,omitempty" yaml:"dir,omitempty"`
	Period    string          `json:"period,omitempty" yaml:"period,omitempty"`
	Expire    int             `json:"expire,omitempty" yaml:"expire,omitempty"`
	Type      string          `json:"type,omitempty" yaml:"type,omitempty"`
	Formatter FormatterConfig `json:"formatter,omitempty" yaml:"formatter,omitempty"`
}
type NsqdOutput struct {
	Config NsqdConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
type NsqdConfig struct {
	Topic      string          `json:"topic,omitempty" yaml:"topic,omitempty"`
	Address    []string        `json:"address,omitempty" yaml:"address,omitempty"`
	ClientConf Config          `json:"nsq_conf,omitempty" yaml:"nsq_conf,omitempty"`
	Type       string          `json:"type,omitempty" yaml:"type,omitempty"`
	Formatter  FormatterConfig `json:"formatter,omitempty" yaml:"formatter,omitempty"`
}
type HttpOutput struct {
	Config HttpConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
type HttpConfig struct {
	Method    string            `json:"method,omitempty" yaml:"method,omitempty"`
	Url       string            `json:"url,omitempty" yaml:"url,omitempty"`
	Headers   map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Type      string            `json:"type,omitempty" yaml:"type,omitempty"`
	Formatter FormatterConfig   `json:"formatter,omitempty" yaml:"formatter,omitempty"`
}

type SysOutput struct {
	Config SysConfig `json:"config,omitempty" yaml:"config,omitempty"`
}
type SysConfig struct {
	Network   string          `json:"network,omitempty" yaml:"network,omitempty"`
	Address   string          `json:"address,omitempty" yaml:"address,omitempty"`
	Level     string          `json:"level,omitempty" yaml:"level,omitempty"`
	Type      string          `json:"type,omitempty" yaml:"type,omitempty"`
	Formatter FormatterConfig `json:"formatter,omitempty" yaml:"formatter,omitempty"`
}
type KafkaOutput struct {
	Config KafkaConfig `json:"config,omitempty" yaml:"config,omitempty"`
}

type KafkaConfig struct {
	Topic         string          `json:"topic,omitempty" yaml:"topic,omitempty"`
	Address       string          `json:"address,omitempty" yaml:"address,omitempty"`
	Timeout       int             `json:"timeout" yaml:"timeout,omitempty"`
	Version       string          `json:"version,omitempty" yaml:"version,omitempty"`
	PartitionType string          `json:"partition_type,omitempty" yaml:"partition_type,omitempty"`
	Partition     int32           `json:"partition,omitempty" yaml:"partition,omitempty"`
	PartitionKey  string          `json:"partition_key,omitempty" yaml:"partition_key,omitempty"`
	Type          string          `json:"type,omitempty" yaml:"type,omitempty"`
	Formatter     FormatterConfig `json:"formatter,omitempty" yaml:"formatter,omitempty"`
}
