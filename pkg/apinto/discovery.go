package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type discovery struct {
	client Client
	url    string
}

func (d *discovery) Get(ctx context.Context, name string) (*v1.Discovery, error) {
	//TODO implement me
	panic("implement me")
}

func (d *discovery) List(ctx context.Context) (*[]response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (d *discovery) Delete(ctx context.Context, discovery *v1.Discovery) error {
	//TODO implement me
	panic("implement me")
}

func (d *discovery) Update(ctx context.Context, discovery *v1.Discovery) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d *discovery) Create(ctx context.Context, discovery *v1.Discovery) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewDiscovery(client Client) *discovery {
	return &discovery{
		url:    fmt.Sprintf("%s/%s", client.Url(), "discovery"),
		client: client,
	}
}
