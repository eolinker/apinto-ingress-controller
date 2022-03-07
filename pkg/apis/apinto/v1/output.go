package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Output struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []OutputSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type OutputSpec struct {
	FileOutput  []FileOutput  `json:"file_output" yaml:"file_output"`
	Nsqd        []NsqdOutput  `json:"nsqd" yaml:"nsqd"`
	HttpOutput  []HttpOutput  `json:"http_output" yaml:"http_output"`
	SysOutput   []SysOutput   `json:"syslog_output" yaml:"syslog_output"`
	KafkaOutput []KafkaOutput `json:"kafka_output" yaml:"kafka_output"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OutputList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Output `json:"items" yaml:"items"`
}

type FileOutput struct {
	Name   string     `json:"name" yaml:"name"`
	Driver string     `json:"driver" yaml:"driver"`
	Config FileConfig `json:"config" yaml:"config"`
}
type FileConfig struct {
	File      string          `json:"file" yaml:"file"`
	Dir       string          `json:"dir" yaml:"dir"`
	Period    string          `json:"period" yaml:"period"`
	Expire    int             `json:"expire" yaml:"expire"`
	Type      string          `json:"type" yaml:"type"`
	Formatter FormatterConfig `json:"formatter" yaml:"formatter"`
}
type NsqdOutput struct {
	Name   string     `json:"name" yaml:"name"`
	Driver string     `json:"driver" yaml:"driver"`
	Config NsqdConfig `json:"config" yaml:"config"`
}
type NsqdConfig struct {
	Topic      string          `json:"topic" yaml:"topic"`
	Address    []string        `json:"address" yaml:"address"`
	ClientConf Config          `json:"nsq_conf" yaml:"nsq_conf"`
	Type       string          `json:"type" yaml:"type"`
	Formatter  FormatterConfig `json:"formatter" yaml:"formatter"`
}
type HttpOutput struct {
	Name   string     `json:"name" yaml:"name"`
	Driver string     `json:"driver" yaml:"driver"`
	Config HttpConfig `json:"config" yaml:"config"`
}
type HttpConfig struct {
	Method    string            `json:"method" yaml:"method"`
	Url       string            `json:"url" yaml:"url"`
	Headers   map[string]string `json:"headers" yaml:"headers"`
	Type      string            `json:"type" yaml:"type"`
	Formatter FormatterConfig   `json:"formatter" yaml:"formatter"`
}
type SysOutput struct {
	Name   string    `json:"name" yaml:"name"`
	Driver string    `json:"driver" yaml:"driver"`
	Config SysConfig `json:"config" yaml:"config"`
}
type SysConfig struct {
	Network   string          `json:"network" yaml:"network"`
	Address   string          `json:"address" yaml:"address"`
	Level     string          `json:"level" yaml:"level"`
	Type      string          `json:"type" yaml:"type"`
	Formatter FormatterConfig `json:"formatter" yaml:"formatter"`
}
type KafkaOutput struct {
	Name   string      `json:"name" yaml:"name"`
	Driver string      `json:"driver" yaml:"driver"`
	Config KafkaConfig `json:"config" yaml:"config"`
}

type KafkaConfig struct {
	Topic         string          `json:"topic" yaml:"topic"`
	Address       string          `json:"address" yaml:"address"`
	Timeout       int             `json:"timeout" yaml:"timeout"`
	Version       string          `json:"version" yaml:"version"`
	PartitionType string          `json:"partition_type" yaml:"partition_type"`
	Partition     int32           `json:"partition" yaml:"partition"`
	PartitionKey  string          `json:"partition_key" yaml:"partition_key"`
	Type          string          `json:"type" yaml:"type"`
	Formatter     FormatterConfig `json:"formatter" yaml:"formatter"`
}
