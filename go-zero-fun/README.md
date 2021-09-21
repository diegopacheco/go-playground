### Install
``bash
go get -u github.com/tal-tech/go-zero
# for Go 1.15 and earlier
GO111MODULE=on go get -u github.com/tal-tech/go-zero/tools/goctl
```

### Commands used
```bash
goctl api -o greet.api
goctl api go -api greet.api -dir greet 
cd greet
go mod init
go mod tidy
go run greet.go -f etc/greet-api.yaml
```
```bash
curl -i http://localhost:8888/greet/from/you
```
