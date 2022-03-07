package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []ServiceSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type ServiceSpec struct {
	Name        string                  `json:"name" yaml:"name"`
	Driver      string                  `json:"driver" yaml:"driver"`
	Desc        string                  `json:"desc" yaml:"desc"`
	Timeout     int64                   `json:"timeout" yaml:"timeout"`
	Retry       int                     `json:"retry" yaml:"retry"`
	RewriteUrl  string                  `json:"rewrite_url" yaml:"rewrite_url"`
	Scheme      string                  `json:"scheme" yaml:"scheme"`
	ProxyMethod string                  `json:"proxy_method" yaml:"proxy_method"`
	Upstream    string                  `json:"upstream" yaml:"upstream"`
	Anonymous   AnonymousConfig         `json:"anonymous" yaml:"anonymous"`
	Plugins     map[string]PluginConfig `json:"plugins" yaml:"plugins"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Service `json:"items" yaml:"items"`
}

type AnonymousConfig struct {
	Type   string `json:"type"`
	Config string `json:"config"`
}
