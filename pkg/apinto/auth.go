package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type auth struct {
	client Client
	url    string
}

func (a *auth) Get(ctx context.Context, name string) (*v1.Auth, error) {
	//TODO implement me
	panic("implement me")
}

func (a *auth) List(ctx context.Context) (*[]response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (a *auth) Delete(ctx context.Context, auth *v1.Auth) error {
	//TODO implement me
	panic("implement me")
}

func (a *auth) Update(ctx context.Context, auth *v1.Auth) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *auth) Create(ctx context.Context, auth *v1.Auth) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuth(client Client) *auth {
	return &auth{
		url:    fmt.Sprintf("%s/%s", client.Url(), "auth"),
		client: client,
	}
}
