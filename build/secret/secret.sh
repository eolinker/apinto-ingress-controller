#!/usr/bin/env bash
# 生成证书的CN为${webhook_service_name}.${namespace of service}.svc
# 表示为webhook服务可调用的host，用户可根据情况配置。
set -e
#生成CA私钥
openssl genrsa -out ca.key 2048
#生成CA公钥
openssl req -new -x509 -days 365000 -subj "/CN=apinto-admission-server.default.svc" -key ca.key -out ca.crt
#生成服务端私钥
openssl genrsa -out server.pem 1024
openssl rsa -in server.pem -out server.key
#生成签发请求
openssl req -new -subj "/CN=apinto-admission-server.default.svc" -key server.pem -out server.csr
#用ca证书签发服务端证书（含公钥）
openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -days 365000 -out server.crt

#最终生成的文件有ca.crt, ca.key, ca.srl, server.crt, server.csr, server.key, server.pem
#其中server.crt,server.key用于controller的https配置证书，以k8s secret资源的方式进行配置
# ca.crt用于ValidatingWebhook配置文件中ca_bundle， 需要将证书里面的值进行base64加密， cat ca.crt | base64