package apinto

import (
	"context"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
)

type Output interface {
	Get(ctx context.Context, name string) (*v1.Output, error)
	List(ctx context.Context) (*[]v1.Output, error)
	Delete(ctx context.Context, output *v1.Output) error
	Update(ctx context.Context, output *v1.Output) (string, error)
	Create(ctx context.Context, output *v1.Output) (string, error)
}

type output struct {
	client Client
	url    string
}

func (o *output) Get(ctx context.Context, name string) (*v1.Output, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) List(ctx context.Context) (*[]v1.Output, error) {
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
