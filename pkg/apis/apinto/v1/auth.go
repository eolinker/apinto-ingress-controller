package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApintoAuth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AuthSpec `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type AuthSpec struct {
	Name              string      `json:"name" yaml:"name,omitempty"`
	Driver            string      `json:"driver,omitempty" yaml:"driver,omitempty"`
	HideCredentials   bool        `json:"hide_credentials,omitempty" yaml:"hide_credentials,omitempty"`
	Credentials       Credentials `json:"credentials,omitempty" yaml:"credentials"`
	RunOnPreflight    bool        `json:"run_on_preflight,omitempty" yaml:"run_on_preflight,omitempty"`
	SignatureIsBase64 bool        `json:"signature_is_base64,omitempty" yaml:"signature_is_base64,omitempty"`
	User              Users       `json:"user,omitempty" yaml:"user,omitempty"`
	ClaimsToVerify    []string    `json:"claims_to_verify,omitempty" yaml:"claims_to_verify,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ApintoAuthList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Items             []ApintoAuth `json:"items,omitempty" yaml:"items,omitempty"`
}

type Users []User

type User struct {
	UserName string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	ApiKey   string `json:"apikey,omitempty" yaml:"apikey,omitempty"`
	Expire   int64  `json:"expire,omitempty" yaml:"expire,omitempty"`
	Ak       string `json:"ak,omitempty" yaml:"ak,omitempty"`
	Sk       string `json:"sk,omitempty" yaml:"sk,omitempty"`
}

type Credentials []Credential

type Credential struct {
	Iss          string `json:"iss,omitempty" yaml:"iss,omitempty"`
	Secret       string `json:"secret,omitempty" yaml:"secret,omitempty"`
	Algorithm    string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
	RsaPublicKey string `json:"rsa_public_key,omitempty" yaml:"rsa_public_key,omitempty"`
}
