/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/kube/apinto/client/clientset/versioned/typed/apinto/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeApintoV1 struct {
	*testing.Fake
}

func (c *FakeApintoV1) ApintoAuths(namespace string) v1.ApintoAuthInterface {
	return &FakeApintoAuths{c, namespace}
}

func (c *FakeApintoV1) ApintoDiscoveries(namespace string) v1.ApintoDiscoveryInterface {
	return &FakeApintoDiscoveries{c, namespace}
}

func (c *FakeApintoV1) ApintoOutputs(namespace string) v1.ApintoOutputInterface {
	return &FakeApintoOutputs{c, namespace}
}

func (c *FakeApintoV1) ApintoRouters(namespace string) v1.ApintoRouterInterface {
	return &FakeApintoRouters{c, namespace}
}

func (c *FakeApintoV1) ApintoServices(namespace string) v1.ApintoServiceInterface {
	return &FakeApintoServices{c, namespace}
}

func (c *FakeApintoV1) ApintoSettings(namespace string) v1.ApintoSettingInterface {
	return &FakeApintoSettings{c, namespace}
}

func (c *FakeApintoV1) ApintoUpstreams(namespace string) v1.ApintoUpstreamInterface {
	return &FakeApintoUpstreams{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeApintoV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
