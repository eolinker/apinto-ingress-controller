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
	seriveListUrl       = ""
	errNotApintoService = errors.New("object is not ApintoService")
)

func SetServiceListUrl(baseURl string) {
	seriveListUrl = baseURl + "/api/service"
}

var ApintoServiceValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成service
		as, ok := object.(*v1.Service)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoService.Error()}, nil
		}
		serviceList := as.Spec

		switch review.Operation {
		case "create", "update":
			//TODO 检查有没有使用upstream，若有，检查upstream是否存在

			objList, err := sourceClient.GetResource(ctx, upstreamListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get UpstreamList fail. " + err.Error()}, nil
			}

			//将所有服务记录到map中
			upstreamSet := make(map[string]struct{})
			for _, serviceJson := range *objList {
				//TODO 反序列化 并且将service id存入serviceSet中
				var upstream v1.Upstream
				err = json.Unmarshal(serviceJson, &upstream)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal Service fail. ServiceJson: %s. err: %s", serviceJson, err)}, nil
				}

				//TODO 定义返回体
				upstreamSet[upstream.Name] = struct{}{}
			}

			//TODO 检查所配置的service是否有效
			for _, service := range serviceList {
				upstreamID := service.Upstream
				if _, has := upstreamSet[upstreamID]; upstreamID != "" && !has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("router:%s config error. upstream:%s not exist. ", service.Name, upstreamID))
				}
			}
		case "delete":
			//TODO 检查有没有router在使用当前服务
			objList, err := sourceClient.GetResource(ctx, routerListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get RouterList fail. " + err.Error()}, nil
			}

			//将对应服务下的所有路由记录到map中
			srMap := make(map[string][]string)
			for _, routerJson := range *objList {
				var router v1.Router
				err = json.Unmarshal(routerJson, &router)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal router fail. RouterJson: %s. err: %s", routerJson, err)}, nil
				}

				//TODO 定义返回体 srMap[router.target] = append(srMap[router.target], router.name)
				srMap[router.Name] = append(srMap[router.Name], router.Name)
			}

			//TODO 检查所配置的service是否有效
			for _, service := range serviceList {
				serviceID := fmt.Sprintf("%s@service", service.Name)
				if l, has := srMap[serviceID]; has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("service:%s is in use by router:%s. ", serviceID, l))
				}
			}

		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
