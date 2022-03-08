package api

import (
	"crypto/tls"
	"github.com/eolinker/apinto-ingress-controller/pkg/api/router"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAdmissionServer(cfg *config.Config) (*http.Server, error) {
	cert, err := tls.LoadX509KeyPair(cfg.CertFilePath, cfg.KeyFilePath)
	if err != nil {
		//TODO 返回报错

		return nil, err
	} else {
		admission := gin.New()
		admission.Use(gin.Recovery(), gin.Logger())
		router.ValidatingWebhook(admission, cfg.APINTO)

		admissionServer := &http.Server{
			Addr:    cfg.HTTPSListen,
			Handler: admission,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}

		return admissionServer, nil
	}

}
