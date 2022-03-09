package validation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	v1 "github.com/eolinker/apinto-ingress-controller/pkg/apis/apinto/v1"
	kwhmodel "github.com/slok/kubewebhook/v2/pkg/model"
	kwhvalidating "github.com/slok/kubewebhook/v2/pkg/webhook/validating"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

var (
	upstreamListUrl      = ""
	errNotApintoUpstream = errors.New("object is not ApintoUpstream")
)

func SetUpstreamListUrl(baseURl string) {
	upstreamListUrl = baseURl + "/api/upstream"
}

var ApintoUpstreamValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成upstream
		au, ok := object.(*v1.Upstream)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoUpstream.Error()}, nil
		}
		upstreamList := au.Spec

		switch review.Operation {
		case "create", "update":
			//TODO 检查使用的discovery是否存在

			objList, err := sourceClient.GetResource(ctx, discoveryListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get DiscoveryList fail. " + err.Error()}, nil
			}

			//将所有discovery记录到map中
			discoverySet := make(map[string]struct{})
			for _, discoveryJson := range *objList {
				//TODO 反序列化 并且将service id存入serviceSet中
				var discovery v1.Discovery
				err = json.Unmarshal(discoveryJson, &discovery)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal Service fail. ServiceJson: %s. err: %s", discoveryJson, err)}, nil
				}

				//TODO 定义返回体
				discoverySet[discovery.Name] = struct{}{}
			}

			//TODO 检查所配置的service是否有效
			for _, upstream := range upstreamList {
				discoveryID := upstream.Discover
				if _, has := discoverySet[discoveryID]; !has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("upstream:%s config error. discovery:%s not exist. ", upstream.Name, discoveryID))
				}
			}
		case "delete":
			//TODO 检查是否有服务在使用这个负载
			objList, err := sourceClient.GetResource(ctx, seriveListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get ServiceList fail. " + err.Error()}, nil
			}

			//将对应负载的所有服务记录到map中
			usMap := make(map[string][]string)
			for _, routerJson := range *objList {
				//TODO 反序列化 并且将service id存入serviceSet中
				var service v1.Service
				err = json.Unmarshal(routerJson, &service)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal service fail. RouterJson: %s. err: %s", routerJson, err)}, nil
				}

				//TODO 定义返回体 usMap[service.upstream] = append(srMap[service.upstream], service.name)
				usMap[service.Name] = append(usMap[service.Name], service.Name)
			}

			//TODO 检查所配置的service是否有效
			for _, upstream := range upstreamList {
				upstreamID := fmt.Sprintf("%s@upstream", upstream.Name)
				if l, has := usMap[upstreamID]; has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("upstream:%s is in use by service:%s. ", upstreamID, l))
				}
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
