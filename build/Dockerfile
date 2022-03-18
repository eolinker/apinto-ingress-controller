FROM golang:1.17-alpine AS builder
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

ENV APPNAME=apinto-ingress-controller

COPY . /go/src/${APPNAME}

RUN cd /go/src/${APPNAME}
RUN go install ./...


FROM alpine:latest
COPY --from=builder /go/src/${APPNAME}/config/config.yaml /etc/ingress/config.yaml
COPY --from=builder /go/bin/${APPNAME} /ingress/${APPNAME}
WORKDIR /ingress

EXPOSE 8080 8443
VOLUME /etc/ingress/

CMD ["/ingress/apinto-ingress-controller", "ingress", "--config-path", "/etc/ingress/config.yaml"]
