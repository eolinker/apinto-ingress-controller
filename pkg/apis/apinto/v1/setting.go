package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Setting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SettingSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type SettingSpec struct {
	Plugins []SettingPlugins `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SettingList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Setting `json:"items,omitempty" yaml:"items,omitempty"`
}

type SettingPlugins []SettingPlugin

type SettingPlugin struct {
	ID     string `json:"id,omitempty" yaml:"id,omitempty"`
	Name   string `json:"name,omitempty" yaml:"name,omitempty"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
	Config Config `json:"config,omitempty" yaml:"config,omitempty"`
}
