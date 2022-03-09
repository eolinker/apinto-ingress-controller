package apinto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type service struct {
	client Client
	url    string
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	url := s.url + "/" + name
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res *v1.Service
	err = json.Unmarshal(*resp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *service) List(ctx context.Context) (*[]response.Response, error) {
	return s.client.List(ctx, s.url)
}

func (s *service) Delete(ctx context.Context, name string) error {
	url := s.url + "/" + name
	_, err := s.client.Delete(ctx, url)
	return err
}

func (s *service) Update(ctx context.Context, service *v1.Service) (string, error) {
	data, err := json.Marshal(service)
	if err != nil {
		return "", err
	}
	url := s.url + "/" + service.Name
	resp, err := s.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (s *service) Create(ctx context.Context, service *v1.Service) (string, error) {
	data, err := json.Marshal(service)
	if err != nil {
		return "", err
	}
	resp, err := s.client.Create(ctx, s.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func NewService(client Client) *service {
	return &service{
		url:    fmt.Sprintf("%s/%s", client.Url(), "service"),
		client: client,
	}
}
