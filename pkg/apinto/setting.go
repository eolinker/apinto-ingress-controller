package apinto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

type setting struct {
	client Client
	url    string
}

func NewSetting(client Client) *setting {
	return &setting{
		url:    fmt.Sprintf("%s/%s", client.Url(), "setting/plugin"),
		client: client,
	}
}

func (s *setting) GetPlugin(ctx context.Context) (*v1.Setting, error) {
	url := s.url
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res *v1.Setting
	err = json.Unmarshal(*resp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *setting) UpdatePlugin(ctx context.Context, setting *v1.Setting) (string, error) {
	data, err := json.Marshal(setting)
	if err != nil {
		return "", err
	}
	resp, err := s.client.Create(ctx, s.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}
