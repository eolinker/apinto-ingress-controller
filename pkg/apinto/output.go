package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type output struct {
	client Client
	url    string
}

func (o *output) Get(ctx context.Context, name string) (*v1.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) List(ctx context.Context) (*[]response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) Delete(ctx context.Context, output *v1.Output) error {
	//TODO implement me
	panic("implement me")
}

func (o *output) Update(ctx context.Context, output *v1.Output) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) Create(ctx context.Context, output *v1.Output) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewOutput(client Client) *output {
	return &output{
		url:    fmt.Sprintf("%s/%s", client.Url(), "output"),
		client: client,
	}
}
