package apinto

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
)

type auth struct {
	client Client
	url    string
}

func (a *auth) Get(ctx context.Context, name string) (*v1.Auth, error) {
	// 先查缓存
	url := a.url + "/" + name
	resp, err := a.client.Get(ctx, url)
	if err != nil {
		return nil, err
	}
	var res *v1.Auth
	err = json.Unmarshal(*resp, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *auth) List(ctx context.Context) (*[]response.Response, error) {
	return a.client.List(ctx, a.url)
}

func (a *auth) Delete(ctx context.Context, name string) error {
	url := a.url + "/" + name
	_, err := a.client.Delete(ctx, url)
	return err
}

func (a *auth) Update(ctx context.Context, auth *v1.Auth) (string, error) {
	data, err := json.Marshal(auth)
	if err != nil {
		return "", err
	}
	url := a.url + "/" + auth.Name
	resp, err := a.client.Update(ctx, url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (a *auth) Create(ctx context.Context, auth *v1.Auth) (string, error) {
	data, err := json.Marshal(auth)
	if err != nil {
		return "", err
	}
	resp, err := a.client.Create(ctx, a.url, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func NewAuth(client Client) *auth {
	return &auth{
		url:    fmt.Sprintf("%s/%s", client.Url(), "auth"),
		client: client,
	}
}
