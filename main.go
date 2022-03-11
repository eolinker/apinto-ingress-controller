package main

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto"
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"github.com/eolinker/eosc/log"
)

func main() {
	client := apinto.NewApinto()
	err := client.AddCluster(&cluster.ClusterOptions{
		Name:    "default",
		BaseURL: "/api/",
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = client.Cluster("test").AuthChecker().DelCheck("")
	if err != nil {
		log.Fatal(err)
	}
}
