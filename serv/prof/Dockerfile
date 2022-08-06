FROM golang:1.17-alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add graphviz
RUN  go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/
RUN mkdir storage
RUN mkdir logs
RUN go install github.com/matteo-gz/prof/cmd/prof@latest
COPY env.yaml /go/src/
EXPOSE 8201
