FROM golang:1.16-alpine
#cn network
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
# go path
RUN PATH=$GOPATH/bin:$PATH
RUN export  GOPATH PATH
RUN  go env -w GO111MODULE=on
RUN  go env -w GOPROXY=https://goproxy.cn,direct
# http 8081
COPY ./demo.go /
# pkg
RUN apk add htop curl graphviz git
#grpc latest
#RUN apk add protoc
#RUN apk add protobuf
RUN apk add  protobuf-dev wget
#grpc  version 3.7.1
ENV PROTOBUF_VERSION="3.7.1"
RUN  if ! [ -e cache/protoc-${PROTOBUF_VERSION}-$(uname -s)-$(uname -m).zip ]; then  wget "https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-$(uname -s)-$(uname -m).zip" -P /tmp/; fi
RUN apk add unzip
RUN unzip /tmp/protoc-${PROTOBUF_VERSION}-$(uname -s)-$(uname -m).zip -d /usr/local
#glibc
ENV GLIBC_VERSION=2.30-r0
RUN  apk --no-cache add wget
RUN wget -q https://alpine-pkgs.sgerrand.com/sgerrand.rsa.pub -O /etc/apk/keys/sgerrand.rsa.pub
RUN wget -q https://github.com/sgerrand/alpine-pkg-glibc/releases/download/${GLIBC_VERSION}/glibc-${GLIBC_VERSION}.apk -O glibc.apk
RUN apk add glibc.apk
#grpc
RUN apk add grpc
RUN  apk add musl-dev
RUN apk add gcc
# work dir
WORKDIR /go/src/