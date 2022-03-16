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
package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	apintoDefaultClusterName = "default"
)

// Config contains all config items which are necessary for
// apinto-ingress-controller's running.
type Config struct {
	CertFilePath    string       `json:"cert_file" yaml:"cert_file"`
	KeyFilePath     string       `json:"key_file" yaml:"key_file"`
	Log             LogConfig    `json:"log" yaml:"log"`
	HTTPListen      string       `json:"http_listen" yaml:"http_listen"`
	HTTPSListen     string       `json:"https_listen" yaml:"https_listen"`
	EnableProfiling bool         `json:"enable_profiling" yaml:"enable_profiling"`
	APINTO          APINTOConfig `json:"apinto" yaml:"apinto"`
}

// KubernetesConfig contains all Kubernetes related config items.
type KubernetesConfig struct {
	Kubeconfig        string   `json:"kubeconfig" yaml:"kubeconfig"`
	AppNamespaces     []string `json:"app_namespaces" yaml:"app_namespaces"`
	NamespaceSelector []string `json:"namespace_selector" yaml:"namespace_selector"`
	IngressClass      string   `json:"ingress_class" yaml:"ingress_class"`
	IngressVersion    string   `json:"ingress_version" yaml:"ingress_version"`
}
type LogConfig struct {
	LogOutput string `json:"log_output" yaml:"log_output"`
	LogLevel  string `json:"log_level" yaml:"log_level"`
	LogPeriod string `json:"log_period" yaml:"log_period"`
	LogExpire string `json:"log_expire" yaml:"log_expire"`
}

// APINTOConfig contains all APINTO related config items.
type APINTOConfig struct {
	DefaultClusterName string `json:"default_cluster_name"`
	// DefaultClusterBaseURL is the base url configuration for the default cluster.
	DefaultClusterBaseURL string `json:"default_cluster_base_url" yaml:"default_cluster_base_url"`
	// DefaultClusterAdminKey is the admin key for the default cluster.
	// TODO: Obsolete the plain way to specify admin_key, which is insecure.
	DefaultClusterAdminKey string `json:"default_cluster_admin_key" yaml:"default_cluster_admin_key"`
}

// NewDefaultConfig creates a Config object which fills all config items with
// default value.
func NewDefaultConfig() *Config {
	return &Config{
		Log: LogConfig{
			LogLevel:  "warn",
			LogOutput: "stderr",
			LogExpire: "1",
			LogPeriod: "day",
		},

		HTTPListen:  ":8080",
		HTTPSListen: ":8443",
		APINTO: APINTOConfig{
			DefaultClusterBaseURL:  "http://127.0.0.1:9400",
			DefaultClusterAdminKey: "",
			DefaultClusterName:     "default",
		},
		EnableProfiling: true,
	}
}

// NewConfigFromFile creates a Config object and fills all config items according
// to the configuration file. The file can be in JSON/YAML format, which will be
// distinguished according to the file suffix.
func NewConfigFromFile(filename string) (*Config, error) {
	cfg := NewDefaultConfig()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	envVarMap := map[string]string{}
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		envVarMap[pair[0]] = pair[1]
	}

	tpl := template.New("text").Option("missingkey=error")
	tpl, err = tpl.Parse(string(data))
	if err != nil {
		return nil, fmt.Errorf("error parsing configuration template %v", err)
	}
	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, envVarMap)
	if err != nil {
		return nil, fmt.Errorf("error execute configuration template %v", err)
	}

	if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		err = yaml.Unmarshal(buf.Bytes(), cfg)
	} else {
		err = json.Unmarshal(buf.Bytes(), cfg)
	}

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// Validate validates whether the Config is right.
func (cfg *Config) Validate() error {

	if cfg.APINTO.DefaultClusterBaseURL == "" {
		return errors.New("apisix base url is required")
	}
	if cfg.APINTO.DefaultClusterName == "" {
		cfg.APINTO.DefaultClusterName = apintoDefaultClusterName
	}

	return nil
}
