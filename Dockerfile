FROM golang:1.17-alpine AS builder
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

ENV APPNAME=apinto-ingress-controller

COPY . /go/src/${APPNAME}

WORKDIR /go/src/${APPNAME}
RUN go install /go/src/${APPNAME}


FROM alpine:latest

ENV APPNAME=apinto-ingress-controller
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY --from=builder /go/bin/${APPNAME} /ingress/${APPNAME}
COPY --from=builder /go/src/${APPNAME}/pkg/config/config.yaml /etc/ingress/config.yaml
COPY --from=builder /go/src/${APPNAME}/certs /etc/webhook/certs/

WORKDIR /ingress

EXPOSE 8080 8443
VOLUME /etc/ingress/

CMD ["/ingress/apinto-ingress-controller", "ingress", "--config-path", "/etc/ingress/config.yaml"]
