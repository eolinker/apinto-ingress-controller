# Apinto-Ingress-Controller

​		Apinto-Ingress-Controller是Kubernetes集群中[Apinto网关](https://github.com/eolinker/apinto)的Ingress控制器，使得Apinto能够作为Ingress资源在集群中运作。
​		该控制器的所有配置均参考Apinto网关的配置格式，并以Kubernetes CRDs（自定义资源）实现。支持配置诸如`router`、`service`、`auth`等Apinto已实现的模块，同时也支持配置插件。



## 概况

* [部署](#部署)
* [快速使用](#快速使用)
* [联系我们](#联系我们)
* [关于我们](#关于我们)



## 部署

由于现阶段不支持以**Helm chart**的方式进行部署，因此暂时以手动部署的形式进行说明。

1. [集群内部署apinto集群](https://help.apinto.com/?path=/quick/arrange)

2. 部署[secret](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy/admission/admission_secret.yml)资源，用于配置controller内的admission server证书。可以使用[脚本](https://github.com/eolinker/apinto-ingress-controller/tree/main/build/secret/secret.sh)生成自签证书进行配置。

   ```shell
   kubectl create -f admission_secret.yml
   ```

3. 部署[configmap](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy/configmap/controller_configmap.yml)资源来生成应用所需配置。

   ```shell
   kubectl create -f controller_configmap.yml
   ```

4. 部署[deployment](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy/deployment/controller_deployment.yml)资源生成应用。

   ```shell
   kubectl create -f controller_deployment.yml
   ```

5. 部署[admission服务](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy/admission/admission_servicce.yml)向validatingWebhook提供校验的服务

   ```shell
   kubectl create -f admission_service.yml
   ```

6. 部署[ValidatingWebhook](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy/admission/webhook_configuration.yml)资源。

   ```shell
   kubectl create -f webhook_configuration.yml
   ```



以上部署所需yaml文件均在[此目录](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/deploy)。

## 快速使用

自定义资源的配置顺序与apinto一致，均有依赖关系。

以配置路由以及服务为例：

**服务 ** `service.yml`

```yaml
apiVersion: apinto.com/v1beta
kind: ApintoService
metadata:
  name: demo-anonymous
spec:
  name: demo-anonymous
  driver: http
  desc: "示例服务"
  timeout: 30000
  anonymous:
    type: round-robin
    config: "http://demo-apinto.eolink.com:8280" #该接口返回http调用信息
  retry: 2
  rewrite_url: /
```

**路由** `router.yml`

```yaml
apiVersion: apinto.com/v1beta
kind: ApintoRouter
metadata:
  name: apinto.router
spec:
  name: apinto.router# 路由名称
  listen: 8080    # 监听端口
  driver: http # 驱动
  protocol: http
  method:
    - GET
  rules:      # 规则列表
    - location: "/demo"  # 匹配路径，该示例为前缀匹配，即只要前缀是 “/” 的路径都可匹配成功
  target: demo-anonymous@service    # 目标服务ID，格式为：{服务名称}@service
```

```shell
kubectl create -f service.yml
kubectl create -f router.yml
```

创建完服务以及路由之后，调用apinto暴露到集群外的服务来查看是否存在该路由

```shell
curl -X GET 'http://{node_ip}:{admin_port}/api/router/apinto.router'
```

返回

```json
{
	"create": "2022-03-23 06:14:48",
	"driver": "http",
	"id": "apinto.router@router",
	"listen": 8080,
	"method": ["GET"],
	"name": "apinto.router",
	"profession": "router",
	"protocol": "http",
	"rules": [{
		"location": "/demo"
	}],
	"target": "demo-anonymous@service",
	"update": "2022-03-23 06:14:48"
}
```

通过调用apinto暴露到集群外的服务来请求该路由

```shell
curl -X GET 'http://{node_ip}:{http_port}/demo'
```

返回

```json
{
	"body": "",
	"header": {
		"Accept": ["*/*"],
		"User-Agent": ["curl/7.75.0"],
		"X-Forwarded-For": ["10.24.1.1,10.24.1.1"]
	},
	"host": "192.2.9.43:31080",  //非原始数据
	"method": "GET",
	"path": "/demo",
	"remote_addr": "192.4.5.22:19091", //非原始数据
	"url": "/demo"
}
```



使用示例[点此](https://github.com/eolinker/apinto-ingress-controller/tree/main/samples/crd/v1/instance)进行跳转。

## 联系我们

- **QQ群**: 725853895
- **Slack**：[加入我们](https://join.slack.com/t/slack-zer6755/shared_invite/zt-u7wzqp1u-aNA0XK9Bdb3kOpN03jRmYQ)
- **官网**：[https://www.apinto.com](https://www.apinto.com/)
- **论坛**：[https://community.apinto.com](https://community.apinto.com/)
- **微信群**：<img src="https://user-images.githubusercontent.com/25589530/149860447-5879437b-3cda-4833-aee3-69a2e538e85d.png" style="width:150px" />

## 关于我们

EOLINK 是领先的 API 管理服务供应商，为全球超过3000家企业提供专业的 API 研发管理、API自动化测试、API监控、API网关等服务。是首家为ITSS（中国电子工业标准化技术协会）制定API研发管理行业规范的企业。

官方网站：[https://www.eolink.com](https://www.eolink.com "EOLINKER官方网站")

免费下载PC桌面端：[https://www.eolink.com/pc/](https://www.eolink.com/pc/ "免费下载PC客户端")

