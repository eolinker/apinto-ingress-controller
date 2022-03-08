package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
)

type Router interface {
	Get(ctx context.Context, name string) (*v1.Router, error)
	List(ctx context.Context) (*[]v1.Router, error)
	Delete(ctx context.Context, name *v1.Router) error
	// Update 反馈id
	Update(ctx context.Context, router *v1.Router) (string, error)
	Create(ctx context.Context, router *v1.Router) (string, error)
}

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

func (r *router) List(ctx context.Context) (*[]v1.Router, error) {
	//TODO implement me
	panic("implement me")
}

func (r *router) Delete(ctx context.Context, name *v1.Router) error {
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
