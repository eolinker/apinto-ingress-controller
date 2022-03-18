package discovery

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type Discovery struct {
	client client.Client
	url    string
}

func (d *Discovery) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Discovery) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Discovery) Get(ctx context.Context, name string) (*v1.Discovery, error) {
	// 先查缓存
	url := d.url + "/" + name
	resp, err := d.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res v1.Discovery
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (d *Discovery) List(ctx context.Context) ([]*response.Response, error) {
	return d.client.List(ctx, d.url)
}

func (d *Discovery) Delete(ctx context.Context, name string) error {
	url := d.url + "/" + name
	_, err := d.client.Delete(ctx, url)
	return err
}

func (d *Discovery) Update(ctx context.Context, discovery *v1.Discovery) (string, error) {
	data, err := json.Marshal(discovery)
	if err != nil {
		return "", err
	}
	url := d.url + "/" + discovery.Name
	resp, err := d.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (d *Discovery) Create(ctx context.Context, discovery *v1.Discovery) (string, error) {
	data, err := json.Marshal(discovery)
	if err != nil {
		return "", err
	}
	resp, err := d.client.Create(ctx, d.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func NewDiscovery(client client.Client) *Discovery {
	return &Discovery{
		url:    fmt.Sprintf("%s/%s", client.Url(), "discovery"),
		client: client,
	}
}
