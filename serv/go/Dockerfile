ARG GO_VERSION
FROM golang:${GO_VERSION}-alpine
#cn network
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
# go path
RUN PATH=$GOPATH/bin:$PATH
RUN export  GOPATH PATH
RUN  go env -w GO111MODULE=on
RUN  go env -w GOPROXY=https://goproxy.cn,direct
RUN apk add htop curl graphviz git
RUN apk add protoc
RUN apk add grpc
RUN apk add  protobuf-dev wget
WORKDIR /go/src/
Env TEST_SERVER=4
EXPOSE 30099
RUN apk add bash

ARG CONFIG_NACOS_HOST
ARG CONFIG_NACOS_USERNAME
ARG CONFIG_NACOS_PASSWD
ARG CONFIG_NACOS_LOGDIR
ARG CONFIG_NACOS_CACHEDIR
ARG CONFIG_NACOS_GROUP

ENV CONFIG_NACOS_HOST=${CONFIG_NACOS_HOST}
ENV CONFIG_NACOS_USERNAME=${CONFIG_NACOS_USERNAME}
ENV CONFIG_NACOS_PASSWD=${CONFIG_NACOS_PASSWD}
ENV CONFIG_NACOS_LOGDIR=${CONFIG_NACOS_LOGDIR}
ENV CONFIG_NACOS_CACHEDIR=${CONFIG_NACOS_CACHEDIR}
ENV CONFIG_NACOS_GROUP=${CONFIG_NACOS_GROUP}
RUN apk add gcc
RUN CGO_ENABLED=0 GOPROXY=https://goproxy.cn/,direct go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.46.2
RUN apk add make