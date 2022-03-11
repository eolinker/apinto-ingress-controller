package kube

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/config"
	clientset "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/clientset/versioned"
	"github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/informers/externalversions"
	"k8s.io/client-go/informers"
	"time"

	"k8s.io/client-go/kubernetes"
)

type Client struct {
	cfg    *config.Config
	Client kubernetes.Interface
	Apinto clientset.Interface
}

func NewClient(cfg *config.Config) (*Client, error) {
	conf, err := BuildRestConfig(cfg.Kubernetes.Kubeconfig, "")
	if err != nil {
		return nil, err
	}
	// 建立k8s连接
	kubeClient, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}
	apinto, err := clientset.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return &Client{
		cfg:    cfg,
		Client: kubeClient,
		Apinto: apinto,
	}, nil
}

// NewSharedIndexInformerFactory 返回一个索引通知器工厂对象，用于监视和列出 Kubernetes 内置资源。
// NewSharedIndexInformerFactory returns an index informer factory object used to watch and
// list Kubernetes builtin resources.
func (k *Client) NewSharedIndexInformerFactory() informers.SharedInformerFactory {
	return informers.NewSharedInformerFactory(k.Client, time.Second*30)
}

// NewAPINTOSharedIndexInformerFactory 返回一个索引通知器工厂对象，用于监视和列出 apisix.apache.org 组中的 Kubernetes 资源。
// NewAPINTOSharedIndexInformerFactory returns an index informer factory object used to watch
// and list Kubernetes resources in apisix.apache.org group.
func (k *Client) NewAPINTOSharedIndexInformerFactory() externalversions.SharedInformerFactory {
	return externalversions.NewSharedInformerFactory(k.Apinto, time.Second*30)
}
