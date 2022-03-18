package ingress

import (
	"context"
	"github.com/eolinker/apinto-ingress-controller/pkg/api"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	scheme2 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/clientset/versioned/scheme"
	"github.com/eolinker/apinto-ingress-controller/pkg/types"
	"github.com/eolinker/eosc/log"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	"net/http"
	"os"
	"time"
)

type Controller struct {
	name      string
	namespace string
	cfg       *config.Config

	// apinto 客户端
	apinto apinto.Apinto
	// admission 客户端，类型待定
	admission *http.Server
}

func NewController(cfg *config.Config) (*Controller, error) {
	podName := os.Getenv("POD_NAME")
	podNamespace := os.Getenv("POD_NAMESPACE")
	if podNamespace == "" {
		podNamespace = "default"
	}
	admission, err := api.NewAdmissionServer(cfg)
	if err != nil {
		return nil, err
	}
	utilruntime.Must(scheme2.AddToScheme(scheme.Scheme))
	return &Controller{
		name:      podName,
		namespace: podNamespace,
		cfg:       cfg,
		admission: admission,
		apinto:    apinto.NewApinto(),
	}, nil
}

// Run 启动controller
func (c *Controller) Run(stopCh <-chan struct{}) error {
	go func() {
		<-stopCh
		closed := make(chan struct{}, 1)
		go c.closeAdmissionServer(closed)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		cnt := 1
		select {
		case <-ctx.Done():
			log.Errorf("close servers timeout")
			return
		case <-closed:
			cnt--
			log.Debug("close a server")
		}
	}()

	// apinto默认集群
	option := &cluster.ClusterOptions{
		Name:     c.cfg.APINTO.DefaultClusterName,
		AdminKey: c.cfg.APINTO.DefaultClusterAdminKey,
		BaseURL:  c.cfg.APINTO.DefaultClusterBaseURL,
	}

	err := c.apinto.AddCluster(option)
	if err != nil {
		log.Errorf("failed to add %s cluster: %s", option.Name, err)
		return err
	}

	// admission监听堵塞
	log.Debug("starting admission server")
	if err := c.admission.ListenAndServeTLS("", ""); err != nil && !types.IsUseOfClosedNetConnErr(err) {
		log.Errorf("failed to start admission server: %s", err)
	}
	return nil
}

func (c *Controller) closeAdmissionServer(closed chan struct{}) {
	if c.admission != nil {
		if err := c.admission.Shutdown(context.TODO()); err != nil {
			log.Errorf("failed to shutdown admission server: %s", err)
		}
	}
	closed <- struct{}{}
}
