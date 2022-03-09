package v1

import "encoding/json"

type PluginConfig struct {
	Disable bool        `json:"disable"`
	Config  interface{} `json:"config"`
}

func (p *PluginConfig) DeepCopyInto(out *PluginConfig) {
	b, _ := json.Marshal(&p)
	_ = json.Unmarshal(b, out)
}

func (p *PluginConfig) DeepCopy() *PluginConfig {
	if p == nil {
		return nil
	}
	out := new(PluginConfig)
	p.DeepCopyInto(out)
	return out
}

type Config map[string]interface{}

func (c *Config) DeepCopyInto(out *Config) {
	b, _ := json.Marshal(&c)
	_ = json.Unmarshal(b, out)
}

func (c *Config) DeepCopy() *Config {
	if c == nil {
		return nil
	}
	out := new(Config)
	c.DeepCopyInto(out)
	return out
}

type FormatterConfig map[string][]string

func (c *FormatterConfig) DeepCopyInto(out *FormatterConfig) {
	b, _ := json.Marshal(&c)
	_ = json.Unmarshal(b, out)
}

func (c *FormatterConfig) DeepCopy() *FormatterConfig {
	if c == nil {
		return nil
	}
	out := new(FormatterConfig)
	c.DeepCopyInto(out)
	return out
}
