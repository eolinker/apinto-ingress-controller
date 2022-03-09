package apinto

import (
	"errors"
	"sync"
)

var (
	_errClusterNotExist   = errors.New("cluster not exist")
	_errClusterDuplicated = errors.New("register cluster duplicated")
)

type Apinto interface {
	Cluster(name string) (Cluster, error)
	AddCluster(*ClusterOptions) error
	UpdateCluster(*ClusterOptions) error
	ListClusters() []Cluster
	DeleteCluster(name string) error
}

type apinto struct {
	mu   sync.RWMutex
	data map[string]Cluster
}

func (a *apinto) Cluster(name string) (Cluster, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	c, ok := a.data[name]
	if !ok {
		return nil, _errClusterNotExist
	}
	return c, nil
}

func (a *apinto) AddCluster(options *ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	_, ok := a.data[options.Name]
	if ok {
		return _errClusterDuplicated
	}
	c, err := NewCluster(options)
	if err != nil {
		return err
	}
	a.data[options.Name] = c
	return nil
}

func (a *apinto) UpdateCluster(options *ClusterOptions) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, ok := a.data[options.Name]; !ok {
		return _errClusterNotExist
	}
	c, err := NewCluster(options)
	if err != nil {
		return err
	}

	a.data[options.Name] = c
	return nil
}

func (a *apinto) ListClusters() []Cluster {
	a.mu.RLock()
	defer a.mu.RUnlock()
	clusters := make([]Cluster, 0, len(a.data))
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
