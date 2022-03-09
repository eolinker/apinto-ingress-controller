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
	routerListUrl      = ""
	errNotApintoRouter = errors.New("object is not ApintoRouter")
)

func SetRouterListUrl(baseURl string) {
	routerListUrl = baseURl + "/api/router"
}

var ApintoRouterValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {

		valid := true
		var msgs []string

		ar, ok := object.(*v1.Router)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoRouter.Error()}, nil
		}
		routerList := ar.Spec

		switch review.Operation {
		case "create", "update":

			objList, err := sourceClient.GetResource(ctx, seriveListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get ServiceList fail. " + err.Error()}, nil
			}

			//将所有服务记录到map中
			serviceSet := make(map[string]struct{})
			for _, serviceJson := range *objList {
				//TODO 反序列化 并且将service id存入serviceSet中
				var service v1.Service
				err = json.Unmarshal(serviceJson, &service)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal Service fail. ServiceJson: %s. err: %s", serviceJson, err)}, nil
				}

				//TODO 定义返回体
				serviceSet[service.Name] = struct{}{}
			}

			//TODO 检查所配置的service是否有效
			for _, router := range routerList {
				target := router.Target
				if _, has := serviceSet[target]; !has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("router:%s config error. target:%s not exist. ", router.Name, target))
				}
			}

		}

		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
