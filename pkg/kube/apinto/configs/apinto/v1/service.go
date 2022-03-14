package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type ServiceSpec struct {
	Name        string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Driver      string                  `json:"driver,omitempty" yaml:"driver,omitempty"`
	Desc        string                  `json:"desc,omitempty" yaml:"desc,omitempty"`
	Timeout     int64                   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Retry       int                     `json:"retry,omitempty" yaml:"retry,omitempty"`
	RewriteUrl  string                  `json:"rewrite_url,omitempty" yaml:"rewrite_url,omitempty"`
	Scheme      string                  `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	ProxyMethod string                  `json:"proxy_method,omitempty" yaml:"proxy_method,omitempty"`
	Upstream    string                  `json:"upstream,omitempty" yaml:"upstream,omitempty"`
	Anonymous   *AnonymousConfig        `json:"anonymous,omitempty" yaml:"anonymous,omitempty"`
	Plugins     map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoServiceList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoService `json:"items,omitempty" yaml:"items,omitempty"`
}

type AnonymousConfig struct {
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Config string `json:"config,omitempty" yaml:"config,omitempty"`
}
