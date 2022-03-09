package apinto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type upstream struct {
	client Client
	url    string
}

func (u *upstream) Get(ctx context.Context, name string) (*v1.Upstream, error) {
	url := u.url + "/" + name
	resp, err := u.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res *v1.Upstream
	err = json.Unmarshal(*resp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *upstream) List(ctx context.Context) (*[]response.Response, error) {
	return u.client.List(ctx, u.url)
}

func (u *upstream) Delete(ctx context.Context, name string) error {
	url := u.url + "/" + name
	_, err := u.client.Delete(ctx, url)
	return err
}

func (u *upstream) Update(ctx context.Context, upstream *v1.Upstream) (string, error) {
	data, err := json.Marshal(upstream)
	if err != nil {
		return "", err
	}
	url := u.url + "/" + upstream.Name
	resp, err := u.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (u *upstream) Create(ctx context.Context, upstream *v1.Upstream) (string, error) {
	data, err := json.Marshal(upstream)
	if err != nil {
		return "", err
	}
	resp, err := u.client.Create(ctx, u.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func NewUpstream(client Client) *upstream {
	return &upstream{
		url:    fmt.Sprintf("%s/%s", client.Url(), "upstream"),
		client: client,
	}
}
