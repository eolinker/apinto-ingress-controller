package apinto

import (
	"errors"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"sync"
)

var (
	_errClusterNotExist   = errors.New("cluster not exist")
	_errClusterDuplicated = errors.New("register cluster duplicated")
)

type Apinto interface {
	Cluster(name string) cluster.Cluster
	AddCluster(*cluster.ClusterOptions) error
	UpdateCluster(*cluster.ClusterOptions) error
	ListClusters() []cluster.Cluster
	DeleteCluster(name string) error
}

func NewApinto() Apinto {
	cli := &apinto{
		dummy: cluster.NewDummy(),
		data:  make(map[string]cluster.Cluster),
	}
	return cli
}

type apinto struct {
	mu    sync.RWMutex
	dummy cluster.Cluster
	data  map[string]cluster.Cluster
}

func (a *apinto) Cluster(name string) cluster.Cluster {
	a.mu.RLock()
	defer a.mu.RUnlock()
	c, ok := a.data[name]
	if !ok {
		return a.dummy
	}
	return c
}

func (a *apinto) AddCluster(options *cluster.ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	_, ok := a.data[options.Name]
	if ok {
		return _errClusterDuplicated
	}
	c, err := cluster.NewCluster(options)
	if err != nil {
		return err
	}
	a.data[options.Name] = c
	return nil
}

func (a *apinto) UpdateCluster(options *cluster.ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.data[options.Name]; !ok {
		return _errClusterNotExist
	}
	c, err := cluster.NewCluster(options)
	if err != nil {
		return err
	}

	a.data[options.Name] = c
	return nil
}

func (a *apinto) ListClusters() []cluster.Cluster {
	a.mu.RLock()
	defer a.mu.RUnlock()
	clusters := make([]cluster.Cluster, 0, len(a.data))
	for _, c := range a.data {
		clusters = append(clusters, c)
	}
	return clusters
}

func (a *apinto) DeleteCluster(name string) error {
	a.mu.Lock()
	delete(a.data, name)
	a.mu.Unlock()
	return nil
}
