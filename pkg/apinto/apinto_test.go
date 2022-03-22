package apinto

import (
	"github.com/eolinker/apinto-ingress-controller/pkg/apinto/cluster"
	"testing"
)

func TestApinto(t *testing.T) {
	client := NewApinto()
	err := client.AddCluster(&cluster.ClusterOptions{
		Name:    "default",
		BaseURL: "/api/",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Cluster("test").AuthChecker().DelCheck("")
	if err != nil {
		t.Fatal(err)
	}
}
