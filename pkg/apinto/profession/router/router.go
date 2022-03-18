package router

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/client"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type Router struct {
	client client.Client
	url    string
}

func NewRouter(client client.Client) *Router {
	return &Router{
		url:    fmt.Sprintf("%s/%s", client.Url(), "router"),
		client: client,
	}
}

func (r *Router) Get(ctx context.Context, name string) (*v1.Router, error) {
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

func (r *Router) List(ctx context.Context) ([]*response.Response, error) {
	return r.client.List(ctx, r.url)
}

func (r *Router) Delete(ctx context.Context, name string) error {
	url := r.url + "/" + name
	_, err := r.client.Delete(ctx, url)
	return err
}

func (r *Router) Update(ctx context.Context, router *v1.Router) (string, error) {
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

func (r *Router) Create(ctx context.Context, router *v1.Router) (string, error) {
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

func (r *Router) DelCheck(name string) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) UpdateCheck(name string, value interface{}) (*response.Response, error) {
	//TODO implement me
	panic("implement me")
}
