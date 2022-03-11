package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoUpstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UpstreamSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type UpstreamSpec struct {
	Name      string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Driver    string                  `json:"driver,omitempty" yaml:"driver,omitempty"`
	Desc      string                  `json:"desc,omitempty" yaml:"desc,omitempty"`
	Discovery string                  `json:"discovery,omitempty" yaml:"discovery,omitempty"`
	Config    string                  `json:"config,omitempty" yaml:"config,omitempty"`
	Scheme    string                  `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Type      string                  `json:"type,omitempty" yaml:"type,omitempty"`
	Plugins   map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoUpstreamList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoUpstream `json:"items,omitempty" yaml:"items,omitempty"`
}
