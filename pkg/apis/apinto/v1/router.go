package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Router struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []RouterSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type RouterSpec struct {
	Name     string                  `json:"name" yaml:"name"`
	Driver   string                  `json:"driver" yaml:"driver"`
	Listen   int                     `json:"listen" yaml:"listen"`
	Target   string                  `json:"target" yaml:"target"`
	Method   []string                `json:"method" yaml:"method"`
	Host     []string                `json:"host" yaml:"host"`
	Protocol string                  `json:"protocol" yaml:"protocol"`
	Cert     []Cert                  `json:"cert" yaml:"cert"`
	Rules    []Rule                  `json:"rules" yaml:"rules"`
	Plugins  map[string]PluginConfig `json:"plugins" yaml:"plugins"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RouterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Router `json:"items" yaml:"items"`
}

//Cert http路由驱动配置证书Cert结构体
type Cert struct {
	Key string `json:"key"`
	Crt string `json:"crt"`
}

//DriverRule http路由驱动配置Rule结构体
type Rule struct {
	Location string            `json:"location" yaml:"location"`
	Header   map[string]string `json:"header" yaml:"header"`
	Query    map[string]string `json:"query" yaml:"query"`
}
