package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Discovery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []DiscoverySpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type DiscoverySpec struct {
	Name     string       `json:"name" yaml:"name"`
	Driver   string       `json:"driver" yaml:"driver"`
	Desc     string       `json:"desc" yaml:"desc"`
	Scheme   string       `json:"scheme" yaml:"scheme"`
	HealthON bool         `json:"health_on" yaml:"health_on"`
	Config   Config       `json:"config" yaml:"config"`
	Health   HealthConfig `json:"health" yaml:"health"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DiscoveryList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Discovery `json:"items" yaml:"items"`
}

type HealthConfig struct {
	Scheme      string `json:"scheme"`
	Method      string `json:"method"`
	URL         string `json:"url"`
	SuccessCode int    `json:"success_code"`
	Period      int    `json:"period"`
	Timeout     int    `json:"timeout"`
}
