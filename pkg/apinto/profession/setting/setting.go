package setting

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type Setting struct {
	client client.Client
	url    string
}

func (s *Setting) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Setting) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func NewSetting(client client.Client) *Setting {
	return &Setting{
		url:    fmt.Sprintf("%s/%s", client.Url(), "setting/plugin"),
		client: client,
	}
}

func (s *Setting) GetPlugin(ctx context.Context) (*v1.Setting, error) {
	url := s.url
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res v1.Setting
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Setting) UpdatePlugin(ctx context.Context, setting *v1.Setting) (string, error) {
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
