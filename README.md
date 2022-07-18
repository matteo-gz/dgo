#### develop for golang
```
#init 
cp .env.example .env
cp docker-compose.yaml.example docker-compose.yaml
#set up
docker-compose up -d
#entry 
docker exec -ti go sh
```
```
go get -u github.com/msoap/go-carpet &&  ln -s $(go env GOPATH)/bin/go-carpet /usr/local/bin/go-carpet
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.42.0
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/vadimi/grpc-client-cli/cmd/grpc-client-cli@v1.11.0
go get github.com/ktr0731/evans
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
```
# host
127.0.0.1 netdata.local.z portainer.local.z beanstalkd-console.local.z prometheus.local.z grafana.local.z etcdm.local.z hotrod.local.z jaeger.local.z
```