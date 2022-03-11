package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eolinker/apinto-ingress-controller/pkg/types/apinto/v1/response"
	"github.com/eolinker/eosc/log"
	"go.uber.org/multierr"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	url2 "net/url"
	"strings"
	"time"
)

const (
	_defaultTimeout = 5 * time.Second
)

var (
	_errReadOnClosedResBody = errors.New("http: read on closed response body")
	_errNotFound            = errors.New("not found")
	_errStillInUse          = errors.New("still in use")
	_defaultTransport       = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 3 * time.Second,
		}).DialContext,
		ResponseHeaderTimeout: 30 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
)

// Client 发送apinto接口请求
// 发送请求时注意是否有admin key
type Client interface {
	Url() string
	Get(ctx context.Context, url string) ([]byte, error)
	List(ctx context.Context, url string) ([]*response.Response, error)
	Create(ctx context.Context, url string, body io.Reader) (*response.Response, error)
	Delete(ctx context.Context, url string) (*response.Response, error)
	Update(ctx context.Context, url string, body io.Reader) (*response.Response, error)
}

type client struct {
	cli         *http.Client
	baseURL     string
	baseURLHost string
	adminKey    string
}

func NewClient(url string, timeout time.Duration, key string) (*client, error) {
	if url == "" {
		return nil, errors.New("empty base url")
	}
	url = strings.TrimSuffix(url, "/")

	u, err := url2.Parse(url)
	if err != nil {
		return nil, err
	}
	if timeout == time.Duration(0) {
		timeout = _defaultTimeout
	}
	return &client{
		baseURL:     url,
		baseURLHost: u.Host,
		cli: &http.Client{
			Timeout:   timeout,
			Transport: _defaultTransport,
		},
		adminKey: key,
	}, nil
}

// 鉴权加载
func (c *client) applyAuth(req *http.Request) {
	if c.adminKey != "" {
		req.Header.Set("APINTO-API-Key", c.adminKey)
	}
}

// 发送请求
func (c *client) do(req *http.Request) (*http.Response, error) {
	c.applyAuth(req)
	return c.cli.Do(req)
}

// 读取并关闭请求体数据流
// drainBody reads whole data until EOF from r, then close it.
func drainBody(r io.ReadCloser, url string) {
	_, err := io.Copy(ioutil.Discard, r)
	if err != nil {
		if err.Error() != _errReadOnClosedResBody.Error() {

			log.Errorf("failed to drain body (read) from %s, err: %s", url, err.Error())
		}
	}

	if err := r.Close(); err != nil {
		log.Errorf("failed to drain body (close) from %s, err: %s", url, err.Error())
	}
}

// 读取响应内容
func readBody(r io.ReadCloser, url string) string {
	defer func() {
		if err := r.Close(); err != nil {
			log.Errorf("failed to close body from %s, err: %s", url, err.Error())
		}
	}()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Errorf("failed to read body from %s, err: %s", url, err.Error())
		return ""
	}
	return string(data)
}

func (c *client) Url() string {
	return c.baseURL
}
func (c *client) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer drainBody(resp.Body, url)
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, _errNotFound
		} else {
			err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
			err = multierr.Append(err, fmt.Errorf("error message: %s", readBody(resp.Body, url)))
		}
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *client) List(ctx context.Context, url string) ([]*response.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	defer drainBody(resp.Body, url)
	if resp.StatusCode != http.StatusOK {
		err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		err = multierr.Append(err, fmt.Errorf("error message: %s", readBody(resp.Body, url)))
		return nil, err
	}
	var list []*response.Response
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *client) Create(ctx context.Context, url string, body io.Reader) (*response.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	defer drainBody(resp.Body, url)

	if resp.StatusCode != http.StatusOK {
		err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		err = multierr.Append(err, fmt.Errorf("error message: %s", readBody(resp.Body, url)))
		return nil, err
	}

	var cr response.Response
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&cr); err != nil {
		return nil, err
	}
	return &cr, nil
}

func (c *client) Delete(ctx context.Context, url string) (*response.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	defer drainBody(resp.Body, url)
	// 错误处理
	if resp.StatusCode != http.StatusOK {
		message := readBody(resp.Body, url)
		if strings.Contains(message, "requiring") {
			return nil, _errStillInUse
		}
		err = multierr.Append(err, fmt.Errorf("error message: %s", message))
		err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		return nil, err
	}
	var res response.Response
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *client) Update(ctx context.Context, url string, body io.Reader) (*response.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	defer drainBody(resp.Body, url)

	if resp.StatusCode != http.StatusOK {
		err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
		err = multierr.Append(err, fmt.Errorf("error message: %s", readBody(resp.Body, url)))
		return nil, err
	}
	var ur response.Response
	dec := json.NewDecoder(resp.Body)
	if err = dec.Decode(&ur); err != nil {
		return nil, err
	}
	return &ur, nil
}
