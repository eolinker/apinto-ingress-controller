package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type router struct {
	client Client
	url    string
}

func NewRouter(client Client) *router {
	return &router{
		url:    fmt.Sprintf("%s/%s", client.Url(), "router"),
		client: client,
	}
}

func (r *router) Get(ctx context.Context, name string) (*v1.Router, error) {
	//TODO implement me
	panic("implement me")
}

func (r *router) List(ctx context.Context) (*[]response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (r *router) Delete(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (r *router) Update(ctx context.Context, router *v1.Router) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *router) Create(ctx context.Context, router *v1.Router) (string, error) {
	//TODO implement me
	panic("implement me")
}
