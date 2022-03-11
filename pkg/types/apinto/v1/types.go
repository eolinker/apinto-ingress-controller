// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1

import "encoding/json"

// Metadata contains all meta information about resources.
// +k8s:deepcopy-gen=true
type Metadata struct {
	Name       string            `json:"name,omitempty" yaml:"name,omitempty"`
	Profession string            `json:"profession,omitempty" yaml:"profession,omitempty"`
	Driver     string            `json:"driver,omitempty" yaml:"driver,omitempty"`
	ID         string            `json:"id,omitempty" yaml:"id,omitempty"`
	Desc       string            `json:"desc,omitempty" yaml:"desc,omitempty"`
	Labels     map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// +k8s:deepcopy-gen=true
type Router struct {
	Metadata `json:",inline" yaml:",inline"`
	Listen   int                     `json:"listen,omitempty" yaml:"listen,omitempty"`
	Target   string                  `json:"target,omitempty" yaml:"target,omitempty"`
	Method   []string                `json:"method,omitempty" yaml:"method,omitempty"`
	Host     []string                `json:"host,omitempty" yaml:"host,omitempty"`
	Protocol string                  `json:"protocol,omitempty" yaml:"protocol,omitempty"`
	Cert     []Cert                  `json:"cert,omitempty" yaml:"cert,omitempty"`
	Rules    []Rule                  `json:"rules,omitempty" yaml:"rules,omitempty"`
	Plugins  map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen=true
type Upstream struct {
	Metadata  `json:",inline" yaml:",inline"`
	Discovery string                  `json:"discovery,omitempty" yaml:"discovery,omitempty"`
	Config    string                  `json:"config,omitempty" yaml:"config,omitempty"`
	Scheme    string                  `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Type      string                  `json:"type,omitempty" yaml:"type,omitempty"`
	Plugins   map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen=true
type Service struct {
	Metadata    `json:",inline" yaml:",inline"`
	Timeout     int64                   `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Retry       int                     `json:"retry,omitempty" yaml:"retry,omitempty"`
	RewriteUrl  string                  `json:"rewrite_url,omitempty" yaml:"rewrite_url,omitempty"`
	Scheme      string                  `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	ProxyMethod string                  `json:"proxy_method,omitempty" yaml:"proxy_method,omitempty"`
	Upstream    string                  `json:"upstream,omitempty" yaml:"upstream,omitempty"`
	Anonymous   *AnonymousConfig        `json:"anonymous,omitempty" yaml:"anonymous,omitempty"`
	Plugins     map[string]PluginConfig `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

// +k8s:deepcopy-gen=true
type Auth struct {
	Metadata          `json:",inline" yaml:",inline"`
	HideCredentials   bool        `json:"hide_credentials,omitempty" yaml:"hide_credentials,omitempty"`
	Credentials       Credentials `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	RunOnPreflight    bool        `json:"run_on_preflight,omitempty" yaml:"run_on_preflight,omitempty"`
	SignatureIsBase64 bool        `json:"signature_is_base64,omitempty" yaml:"signature_is_base64,omitempty"`
	User              Users       `json:"user,omitempty" yaml:"user,omitempty"`
	ClaimsToVerify    []string    `json:"claims_to_verify,omitempty" yaml:"claims_to_verify,omitempty"`
}

// +k8s:deepcopy-gen=true
type Output struct {
	Metadata `json:",inline" yaml:",inline"`
	Config   Config `json:"config,omitempty" yaml:"config,omitempty"`
}

// +k8s:deepcopy-gen=true
type Discovery struct {
	Metadata `json:",inline" yaml:",inline"`
	Scheme   string       `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	HealthON bool         `json:"health_on,omitempty" yaml:"health_on"`
	Config   Config       `json:"config,omitempty" yaml:"config,omitempty"`
	Health   HealthConfig `json:"health,omitempty" yaml:"health,omitempty"`
}

// +k8s:deepcopy-gen=true
type Setting struct {
	Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
	Profession string         `json:"profession,omitempty" yaml:"profession,omitempty"`
	Driver     string         `json:"driver,omitempty" yaml:"driver,omitempty"`
	Plugins    SettingPlugins `json:"plugins,omitempty" yaml:"plugins,omitempty"`
}

//Cert http路由驱动配置证书Cert结构体
type Cert struct {
	Key string `json:"key,omitempty"`
	Crt string `json:"crt,omitempty"`
}

//DriverRule http路由驱动配置Rule结构体
type Rule struct {
	Location string            `json:"location,omitempty" yaml:"location,omitempty"`
	Header   map[string]string `json:"header,omitempty" yaml:"header,omitempty"`
	Query    map[string]string `json:"query,omitempty" yaml:"query,omitempty"`
}

func (c *Rule) DeepCopyInto(out *Rule) {
	b, _ := json.Marshal(&c)
	_ = json.Unmarshal(b, out)
}

func (c *Rule) DeepCopy() *Rule {
	if c == nil {
		return nil
	}
	out := new(Rule)
	c.DeepCopyInto(out)
	return out
}

type AnonymousConfig struct {
	Type   string `json:"type,omitempty"`
	Config string `json:"config,omitempty"`
}

type SettingPlugins []SettingPlugin

type SettingPlugin struct {
	ID         string `json:"id,omitempty" yaml:"id,omitempty"`
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string `json:"type,omitempty" yaml:"type,omitempty"`
	Status     string `json:"status,omitempty" yaml:"status,omitempty"`
	Config     Config `json:"config,omitempty" yaml:"config,omitempty"`
	InitConfig Config `json:"init_config,omitempty" yaml:"init_config,omitempty"`
}

func (c *SettingPlugin) DeepCopyInto(out *SettingPlugin) {
	b, _ := json.Marshal(&c)
	_ = json.Unmarshal(b, out)
}

func (c *SettingPlugin) DeepCopy() *SettingPlugin {
	if c == nil {
		return nil
	}
	out := new(SettingPlugin)
	c.DeepCopyInto(out)
	return out
}

type HealthConfig struct {
	Scheme      string `json:"scheme,omitempty" yaml:"scheme,omitempty"`
	Method      string `json:"method,omitempty" yaml:"method,omitempty"`
	URL         string `json:"url,omitempty" yaml:"url,omitempty"`
	SuccessCode int    `json:"success_code,omitempty" yaml:"success_code,omitempty"`
	Period      int    `json:"period,omitempty" yaml:"period,omitempty"`
	Timeout     int    `json:"timeout,omitempty" yaml:"timeout,omitempty"`
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

func (c *Credential) DeepCopyInto(out *Credential) {
	b, _ := json.Marshal(&c)
	_ = json.Unmarshal(b, out)
}

func (c *Credential) DeepCopy() *Credential {
	if c == nil {
		return nil
	}
	out := new(Credential)
	c.DeepCopyInto(out)
	return out
}
