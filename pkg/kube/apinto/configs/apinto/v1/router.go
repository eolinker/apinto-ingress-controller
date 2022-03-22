package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouterSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type RouterSpec struct {
	Name     string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Driver   string                  `json:"driver,omitempty" yaml:"driver,omitempty"`
	Listen   int                     `json:"listen,omitempty" yaml:"listen,omitempty"`
	Target   string                  `json:"target,omitempty" yaml:"target,omitempty"`
	Method   []string                `json:"method,omitempty" yaml:"method,omitempty"`
	Host     []string                `json:"host,omitempty" yaml:"host,omitempty"`
	Protocol string                  `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Cert     []Cert                  `json:"cert,omitempty" yaml:"cert,omitempty"`
	Rules    []Rule                  `json:"rules,omitempty" yaml:"rules,omitempty"`
	Plugins  map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoRouterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoRouter `json:"items,omitempty" yaml:"items,omitempty"`
}

//Cert http路由驱动配置证书Cert结构体
type Cert struct {
	Key string `json:"key,omitempty" yaml:"key,omitempty"`
	Crt string `json:"crt,omitempty" yaml:"crt,omitempty"`
}

//DriverRule http路由驱动配置Rule结构体
type Rule struct {
	Location string            `json:"location,omitempty" yaml:"location,omitempty"`
	Header   map[string]string `json:"header,omitempty" yaml:"header,omitempty"`
	Query    map[string]string `json:"query,omitempty" yaml:"query,omitempty"`
}
