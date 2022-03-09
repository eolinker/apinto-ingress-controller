package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

type Service interface {
	Get(ctx context.Context, name string) (*v1.Service, error)
	List(ctx context.Context) (*[]v1.Service, error)
	Delete(ctx context.Context, service *v1.Service) error
	Update(ctx context.Context, service *v1.Service) (string, error)
	Create(ctx context.Context, service *v1.Service) (string, error)
}

type service struct {
	client Client
	url    string
}

func (s *service) Get(ctx context.Context, name string) (*v1.Service, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) List(ctx context.Context) (*[]v1.Service, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Delete(ctx context.Context, service *v1.Service) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) Update(ctx context.Context, service *v1.Service) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Create(ctx context.Context, service *v1.Service) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(client Client) *service {
	return &service{
		url:    fmt.Sprintf("%s/%s", client.Url(), "service"),
		client: client,
	}
}
