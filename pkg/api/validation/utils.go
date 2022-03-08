package validation

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	kwhhttp "github.com/slok/kubewebhook/v2/pkg/http"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	"io"
	"io/ioutil"
	"net/http"
)

type SourceClient struct {
	cli *http.Client
}

type listResponse []json.RawMessage

var (
	_errReadOnClosedResBody = errors.New("http: read on closed response body")
	client                  *SourceClient
)

func NewHandler(ID string, validator kwhvalidating.Validator) gin.HandlerFunc {
	wh, err := kwhvalidating.NewWebhook(kwhvalidating.WebhookConfig{
		ID:        ID,
		Validator: validator,
	})

	if err != nil {
		//TODO 打印日志
		//log.Errorf("failed to create webhook: %s", err)
	}

	h, err := kwhhttp.HandlerFor(kwhhttp.HandlerConfig{Webhook: wh})
	if err != nil {
		//TODO 打印日志
		//	log.Errorf("failed to create webhook handle: %s", err)
	}

	return gin.WrapH(h)
}

func (s *SourceClient) GetResource(ctx context.Context, url string) (*listResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.do(req)
	if err != nil {
		return nil, err
	}

	defer drainBody(resp.Body, url)
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, errors.New("No found. ")
		} else {
			//err = multierr.Append(err, fmt.Errorf("unexpected status code %d", resp.StatusCode))
			//err = multierr.Append(err, fmt.Errorf("error message: %s", readBody(resp.Body, url)))
		}
		return nil, err
	}

	var list listResponse

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (s *SourceClient) do(req *http.Request) (*http.Response, error) {
	//TODO applyAuth(req)
	return s.cli.Do(req)
}

func drainBody(r io.ReadCloser, url string) {
	_, err := io.Copy(ioutil.Discard, r)
	if err != nil {
		if err.Error() != _errReadOnClosedResBody.Error() {
			//log.Warnw("failed to drain body (read)",
			//	zap.String("url", url),
			//	zap.Error(err),
			//)
		}
	}

	if err := r.Close(); err != nil {
		//log.Warnw("failed to drain body (close)",
		//	zap.String("url", url),
		//	zap.Error(err),
		//)
	}
}
