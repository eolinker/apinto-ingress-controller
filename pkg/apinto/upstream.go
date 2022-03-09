package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
)

type Upstream interface {
	Get(ctx context.Context, name string) (*v1.Upstream, error)
	List(ctx context.Context) (*[]v1.Upstream, error)
	Delete(ctx context.Context, upstream *v1.Upstream) error
	Update(ctx context.Context, upstream *v1.Upstream) (string, error)
	Create(ctx context.Context, upstream *v1.Upstream) (string, error)
}

type upstream struct {
	client Client
	url    string
}

func (u *upstream) Get(ctx context.Context, name string) (*v1.Upstream, error) {
	//TODO implement me
	panic("implement me")
}

func (u *upstream) List(ctx context.Context) (*[]v1.Upstream, error) {
	//TODO implement me
	panic("implement me")
}

func (u *upstream) Delete(ctx context.Context, upstream *v1.Upstream) error {
	//TODO implement me
	panic("implement me")
}

func (u *upstream) Update(ctx context.Context, upstream *v1.Upstream) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (u *upstream) Create(ctx context.Context, upstream *v1.Upstream) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewUpstream(client Client) *upstream {
	return &upstream{
		url:    fmt.Sprintf("%s/%s", client.Url(), "upstream"),
		client: client,
	}
}
