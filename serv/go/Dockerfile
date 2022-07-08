FROM golang:1.17-alpine
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