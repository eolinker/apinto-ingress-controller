package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoDiscovery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiscoverySpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type DiscoverySpec struct {
	Name     string       `json:"name,omitempty" yaml:"name,omitempty"`
	Driver   string       `json:"driver,omitempty" yaml:"driver,omitempty"`
	Desc     string       `json:"desc,omitempty" yaml:"desc,omitempty"`
	Scheme   string       `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	HealthON bool         `json:"health_on,omitempty" yaml:"health_on,omitempty"`
	Config   Config       `json:"config,omitempty" yaml:"config,omitempty"`
	Health   HealthConfig `json:"health,omitempty" yaml:"health,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ApintoDiscoveryList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoDiscovery `json:"items,omitempty" yaml:"items,omitempty"`
}

type HealthConfig struct {
	Scheme      string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Method      string `json:"method,omitempty" yaml:"method,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
	SuccessCode int    `json:"success_code,omitempty" yaml:"success_code,omitempty"`
	Period      int    `json:"period,omitempty" yaml:"period,omitempty"`
	Timeout     int    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
