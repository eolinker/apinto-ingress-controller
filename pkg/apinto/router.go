package apinto

import (
	"bytes"
	"context"
	"encoding/json"
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
	// 先查缓存
	url := r.url + "/" + name
	resp, err := r.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res v1.Router
	err = json.Unmarshal(resp, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *router) List(ctx context.Context) ([]*response.Response, error) {
	return r.client.List(ctx, r.url)
}

func (r *router) Delete(ctx context.Context, name string) error {
	url := r.url + "/" + name
	_, err := r.client.Delete(ctx, url)
	return err
}

func (r *router) Update(ctx context.Context, router *v1.Router) (string, error) {
	data, err := json.Marshal(router)
	if err != nil {
		return "", err
	}
	url := r.url + "/" + router.Name
	resp, err := r.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (r *router) Create(ctx context.Context, router *v1.Router) (string, error) {
	data, err := json.Marshal(router)
	if err != nil {
		return "", err
	}
	resp, err := r.client.Create(ctx, r.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (r *router) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (r *router) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}
