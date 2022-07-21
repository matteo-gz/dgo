FROM golang:1.17-alpine
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
RUN apk add graphviz
RUN  go env -w GOPROXY=https://goproxy.cn,direct
RUN go install github.com/matteo-gz/prof/cmd/prof@v1.0.2
WORKDIR /go/src/
RUN cp /go/pkg/mod/github.com/matteo-gz/prof@v1.0.2/web . -r
RUN mkdir storage
RUN mkdir logs
COPY env.yaml /go/src/
EXPOSE 8201