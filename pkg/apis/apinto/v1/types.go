package v1

import (
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Upstream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []UpstreamSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type UpstreamSpec struct {
	Name     string             `json:"name" yaml:"name"`
	Driver   string             `json:"driver" yaml:"driver"`
	Desc     string             `json:"desc" yaml:"desc"`
	Discover string             `json:"discover" yaml:"discover"`
	Config   string             `json:"config" yaml:"config"`
	Scheme   string             `json:"scheme" yaml:"scheme"`
	Type     string             `json:"type" yaml:"type"`
	Plugins  map[string]Plugins `json:"plugins" yaml:"plugins"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type UpstreamList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Upstream `json:"items" yaml:"items"`
}

type Plugins struct {
	Disable bool        `json:"disable"`
	Config  interface{} `json:"config"`
}

func (p *Plugins) DeepCopyInto(out *Plugins) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *Plugins) DeepCopy() *Plugins {
	if p == nil {
		return nil
	}
	out := new(Plugins)
	p.DeepCopyInto(out)
	return out
}
