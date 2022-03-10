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
	discoveryListUrl      = ""
	errNotApintoDiscovery = errors.New("object is not ApintoDiscovery")
)

func SetDiscoveryListUrl(baseURl string) {
	discoveryListUrl = baseURl + "/api/discovery"
}

var ApintoDiscoveryValidator = kwhvalidating.ValidatorFunc(
	func(ctx context.Context, review *kwhmodel.AdmissionReview, object metav1.Object) (result *kwhvalidating.ValidatorResult, err error) {
		valid := true
		var msgs []string

		//将object转化成discovery
		ad, ok := object.(*v1.Discovery)
		if !ok {
			return &kwhvalidating.ValidatorResult{Valid: false, Message: errNotApintoDiscovery.Error()}, nil
		}
		discoveryList := ad.Spec

		switch review.Operation {
		case "delete":
			//TODO 检查是否有upstream在使用这个discovery
			objList, err := sourceClient.GetResource(ctx, upstreamListUrl)
			if err != nil {
				return &kwhvalidating.ValidatorResult{Valid: false, Message: "Get UpstreamList fail. " + err.Error()}, nil
			}

			//将对应discovery下的所有upstream记录到map中
			duMap := make(map[string][]string)
			for _, upstreamJson := range *objList {
				var upstream v1.Upstream
				err = json.Unmarshal(upstreamJson, &upstream)
				if err != nil {
					return &kwhvalidating.ValidatorResult{Valid: false, Message: fmt.Sprintf("unmarshal upstream fail. upstreamJson: %s. err: %s", upstreamJson, err)}, nil
				}

				//TODO 定义返回体 usMap[upstream.upstream] = append(srMap[upstream.upstream], upstream.name)
				duMap[upstream.Name] = append(duMap[upstream.Name], upstream.Name)
			}

			//TODO 检查所配置的service是否有效
			for _, discovery := range discoveryList {
				discoveryID := fmt.Sprintf("%s@discovery", discovery.Name)
				if l, has := duMap[discoveryID]; has {
					valid = false
					msgs = append(msgs, fmt.Sprintf("discovery:%s is in use by upstream:%s. ", discoveryID, l))
				}
			}
		}
		return &kwhvalidating.ValidatorResult{Valid: valid, Message: strings.Join(msgs, "\n")}, nil
	},
)
