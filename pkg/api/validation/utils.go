package validation

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	kwhhttp "github.com/slok/kubewebhook/v2/pkg/http"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	"net/http"
	"time"
)

type SourceClient struct {
	cli *http.Client
}

type listResponse []json.RawMessage

const _defaultTimeout = 5 * time.Second

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
