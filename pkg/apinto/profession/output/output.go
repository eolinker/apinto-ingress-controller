package output

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type Output struct {
	client client.Client
	url    string
}

func (o *Output) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (o *Output) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (o *Output) Get(ctx context.Context, name string) (*v1.Output, error) {
	url := o.url + "/" + name
	resp, err := o.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res v1.Output
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (o *Output) List(ctx context.Context) ([]*response.Response, error) {
	return o.client.List(ctx, o.url)
}

func (o *Output) Delete(ctx context.Context, name string) error {
	url := o.url + "/" + name
	_, err := o.client.Delete(ctx, url)
	return err
}

func (o *Output) Update(ctx context.Context, output *v1.Output) (string, error) {
	data, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	url := o.url + "/" + output.Name
	resp, err := o.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (o *Output) Create(ctx context.Context, output *v1.Output) (string, error) {
	data, err := json.Marshal(output)
	if err != nil {
		return "", err
	}
	resp, err := o.client.Create(ctx, o.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func NewOutput(client client.Client) *Output {
	return &Output{
		url:    fmt.Sprintf("%s/%s", client.Url(), "output"),
		client: client,
	}
}
