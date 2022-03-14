package validation

import (
	"github.com/eolinker/eosc/log"
	"github.com/gin-gonic/gin"
	kwhhttp "github.com/slok/kubewebhook/v2/pkg/http"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	"net/http"
)

type SourceClient struct {
	cli *http.Client
}

func NewHandler(ID string, validator kwhvalidating.Validator) gin.HandlerFunc {
	wh, err := kwhvalidating.NewWebhook(kwhvalidating.WebhookConfig{
		ID:        ID,
		Validator: validator,
	})

	if err != nil {
		log.Errorf("failed to create webhook: %s", err)
	}

	h, err := kwhhttp.HandlerFor(kwhhttp.HandlerConfig{Webhook: wh})
	if err != nil {
		log.Errorf("failed to create webhook: %s", err)
	}

	return gin.WrapH(h)
}
