package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type Service struct {
	client client.Client
	url    string
}

func (s *Service) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Get(ctx context.Context, name string) (*v1.Service, error) {
	url := s.url + "/" + name
	resp, err := s.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res v1.Service
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *Service) List(ctx context.Context) ([]*response.Response, error) {
	return s.client.List(ctx, s.url)
}

func (s *Service) Delete(ctx context.Context, name string) error {
	url := s.url + "/" + name
	_, err := s.client.Delete(ctx, url)
	return err
}

func (s *Service) Update(ctx context.Context, service *v1.Service) (string, error) {
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

func (s *Service) Create(ctx context.Context, service *v1.Service) (string, error) {
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

func NewService(client client.Client) *Service {
	return &Service{
		url:    fmt.Sprintf("%s/%s", client.Url(), "service"),
		client: client,
	}
}
