package api

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/api/router"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAdmissionServer(cfg *config.Config) (*http.Server, error) {
	//cert, err := tls.LoadX509KeyPair(cfg.CertFilePath, cfg.KeyFilePath)
	//if err != nil {
	//	return nil, err
	//} else {
	//	admission := gin.New()
	//	admission.Use(gin.Recovery(), gin.Logger())
	//	err = router.ValidatingWebhook(admission, cfg)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	admissionServer := &http.Server{
	//		Addr:    cfg.HTTPSListen,
	//		Handler: admission,
	//		TLSConfig: &tls.Config{
	//			Certificates: []tls.Certificate{cert},
	//		},
	//	}
	//
	//	return admissionServer, nil
	//}
	admission := gin.New()
	admission.Use(gin.Recovery(), gin.Logger())
	err := router.ValidatingWebhook(admission, cfg)
	if err != nil {
		return nil, err
	}

	admissionServer := &http.Server{
		Addr:    cfg.HTTPListen,
		Handler: admission,
	}

	return admissionServer, nil
}
