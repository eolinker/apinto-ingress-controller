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
	Plugins []SettingPlugins `json:"plugins" yaml:"plugins"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SettingList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Setting `json:"items" yaml:"items"`
}

type SettingPlugins []SettingPlugin

type SettingPlugin struct {
	ID     string `json:"id" yaml:"id"`
	Name   string `json:"name" yaml:"name"`
	Type   string `json:"type" yaml:"type"`
	Status string `json:"status" yaml:"status"`
	Config Config `json:"config" yaml:"config"`
}
