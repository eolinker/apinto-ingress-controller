package router

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/api/validation"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/gin-gonic/gin"
)

func ValidatingWebhook(g *gin.Engine, cfg *config.Config, apinto apinto.Apinto) error {
	err := validation.InitAdmission(cfg, apinto)

	if err != nil {
		return err
	}

	vGroup := g.Group("/validation")
	{
		vGroup.POST("/router", validation.NewHandler("ApintoRouter", validation.ApintoRouterValidator))
		vGroup.POST("/service", validation.NewHandler("ApintoService", validation.ApintoServiceValidator))
		vGroup.POST("/upstream", validation.NewHandler("ApintoUpstream", validation.ApintoUpstreamValidator))
		vGroup.POST("/discovery", validation.NewHandler("ApintoDiscovery", validation.ApintoDiscoveryValidator))
		vGroup.POST("/output", validation.NewHandler("ApintoOutout", validation.ApintoOutputValidator))
		vGroup.POST("/auth", validation.NewHandler("ApintoAuth", validation.ApintoAuthValidator))
		vGroup.POST("/setting", validation.NewHandler("ApintoGlobalSetting", validation.ApintoGlobalSettingValidator))
	}

	return nil
}
