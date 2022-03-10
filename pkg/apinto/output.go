package apinto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type output struct {
	client Client
	url    string
}

func (o *output) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (o *output) Get(ctx context.Context, name string) (*v1.Output, error) {
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

func (o *output) List(ctx context.Context) ([]*response.Response, error) {
	return o.client.List(ctx, o.url)
}

func (o *output) Delete(ctx context.Context, name string) error {
	url := o.url + "/" + name
	_, err := o.client.Delete(ctx, url)
	return err
}

func (o *output) Update(ctx context.Context, output *v1.Output) (string, error) {
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

func (o *output) Create(ctx context.Context, output *v1.Output) (string, error) {
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

func NewOutput(client Client) *output {
	return &output{
		url:    fmt.Sprintf("%s/%s", client.Url(), "output"),
		client: client,
	}
}
