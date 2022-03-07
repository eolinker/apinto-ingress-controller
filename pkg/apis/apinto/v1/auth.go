package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient

// +genclient:noStatus

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Auth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              []AuthSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type AuthSpec struct {
	Name              string      `json:"name" yaml:"name"`
	Driver            string      `json:"driver" yaml:"driver"`
	HideCredentials   bool        `json:"hide_credentials" yaml:"hide_credentials"`
	Credentials       Credentials `json:"credentials" yaml:"credentials"`
	RunOnPreflight    bool        `json:"run_on_preflight" yaml:"run_on_preflight"`
	SignatureIsBase64 bool        `json:"signature_is_base64" yaml:"signature_is_base64"`
	User              Users       `json:"user" yaml:"user"`
	ClaimsToVerify    []string    `json:"claims_to_verify" yaml:"claims_to_verify"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []Auth `json:"items" yaml:"items"`
}

type Users []User

type User struct {
	UserName string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	ApiKey   string `json:"apikey" yaml:"apikey"`
	Expire   int64  `json:"expire" yaml:"expire"`
	Ak       string `json:"ak" yaml:"ak"`
	Sk       string `json:"sk" yaml:"sk"`
}

type Credentials []Credential

type Credential struct {
	Iss          string `json:"iss" yaml:"iss"`
	Secret       string `json:"secret" yaml:"secret"`
	Algorithm    string `json:"algorithm" yaml:"algorithm"`
	RsaPublicKey string `json:"rsa_public_key" yaml:"rsa_public_key"`
}
