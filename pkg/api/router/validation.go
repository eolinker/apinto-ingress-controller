package router

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/api/validation"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/gin-gonic/gin"
)

func ValidatingWebhook(g *gin.Engine, cfg config.APINTOConfig) {
	vGroup := g.Group("/validation")

	{
		vGroup.POST("/apintorouters", validation.NewHandler("ApintoRouter", validation.ApintoRouterValidator))
		vGroup.POST("/apintoservices", validation.NewHandler("ApintoService", validation.ApintoServiceValidator))
		vGroup.POST("/apintoupstreams", validation.NewHandler("ApintoUpstream", validation.ApintoUpstreamValidator))
		vGroup.POST("/apintodiscoveries", validation.NewHandler("ApintoDiscovery", validation.ApintoDiscoveryValidator))
		vGroup.POST("/apintooutputs", validation.NewHandler("ApintoOutout", validation.ApintoOutputValidator))
		vGroup.POST("/apintoauths", validation.NewHandler("ApintoAuth", validation.ApintoAuthValidator))
		vGroup.POST("/apintoglobalsettings", validation.NewHandler("ApintoGlobalSetting", validation.ApintoGlobalSettingValidator))
	}

	validation.SetRouterListUrl(cfg.DefaultClusterBaseURL)
	validation.SetServiceListUrl(cfg.DefaultClusterBaseURL)
	validation.SetUpstreamListUrl(cfg.DefaultClusterBaseURL)
	validation.SetDiscoveryListUrl(cfg.DefaultClusterBaseURL)
	validation.SetAuthListUrl(cfg.DefaultClusterBaseURL)
	validation.SetOutputListUrl(cfg.DefaultClusterBaseURL)
	validation.SetGlobalSettingListUrl(cfg.DefaultClusterBaseURL)

	validation.InitSourceClient()
}
