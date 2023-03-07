# notes

> https://github.com/mingrammer/go-web-framework-stars

## chapter 01

API test:

```bash
go install github.com/narenaryan/romanserver
# check GOPATH: bin/romanserver.exe
```

```http request
GET http://localhost:8000/roman_number/3
"III"
```

```http request
GET http://localhost:8000/roman_number/12
404 - Not Found
```

```http request
GET http://localhost:8000/roman_no_api/12
400 - Bad request
```

supervisor:

```bash
sudo apt-get install -y supervisor

# refer: /etc/supervisor/supervisord.conf
nano /etc/supervisor/conf.d/goproject.conf
```

```bash
root@DESKTOP-QLDBOG2:/etc/supervisor# supervisorctl
supervisor> help

default commands (type help <topic>):
=====================================
add    exit      open  reload  restart   start   tail
avail  fg        pid   remove  shutdown  status  update
clear  maintail  quit  reread  signal    stop    version
```

```bash
# supervisorctl reload
# supervisorctl status
romanserver                      RUNNING   pid 1140, uptime 0:07:00
```

最后的 gulp 自动化不实用，理解即可。

## chapter 02

强大的路由器 1

```bash
go get github.com/julienschmidt/httprouter
```

强大的路由器 2

```bash
go get -u github.com/gorilla/mux
```

## chapter 03

中间件与 RPC。

> Chapter03/cityAPI.go

```bash
curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"New York","area":304}'

curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston","area":89}'
```

> multipleMiddleware.go

```bash
curl -H "Content-Type: application/json" -X POST http://localhost:8000/city -d '{"name":"Boston","area":89}'
```

```
2023/03/05 12:28:58 Currently in the check content type middleware
2023/03/05 12:28:58 Got Boston city with area of 89 sq miles!  
2023/03/05 12:28:58 Currently in the set server time middleware
```

Alice 中间件：简化中间件链式调用的 API 语法。参考：`Chapter03/multipleMiddlewareWithAlice.go`

```bash
go get github.com/justinas/alice
```

Gorilla loggedRouter：内置实用的 logging 日志格式

---
RPC:

自定义 RPC 代码只有在客户端和服务端都用 Go 写的时候才有用（C/S 同构）。为了让 RPC 服务端被多个服务消费，需要定义 JSON RPC over HTTP。
然后，任何其他编程语言都可以发送一个 JSON 字符串并得到 JSON 作为结果。

RPC 面向的是函数调用，区别于 REST 面向的是资源的操作。

> Chapter03/RPCServer.go 代码中注意 port 不要使用 1234，而是更大值的端口，避免权限错误。

JSONRPC:

```bash
curl -X POST http://localhost:12345/rpc \
 -H 'cache-control: no-cache' \
 -H 'content-type: application/json' \
 -d '{
 "method": "JSONServer.GiveBookDetail",
 "params": [{"Id": "1234"}],
 "id": "1"
}'
```

注意请求数据中的这部分：

```
 "params": [{"Id": "1234"}],
 "id": "1"
```

Id 1234 对应 args.Id：

```go
type Args struct {
	Id string
}
```

"id": "1" 意义不明？经过测试，缺失这个 id 不会拿到正确的响应。

```bash
{"result":{"Id":"1234","Name":"In the sunburned country","Author":"Bill Bryson"},"error":null,"id":"1"}
```

## chapter 04

- go‑restful
- gin
- Revel

Make sure your GOPROXY setting value can be reached.

```bash
apt-get install sqlite3 libsqlite3-dev
go get github.com/emicklei/go-restful
go get github.com/mattn/go-sqlite3
```

### gin

```bash
go get gopkg.in/gin-gonic/gin.v1
```

go run `railAPIGin/main.go`:

POST

```bash
$ curl -X POST \
>  http://localhost:8000/v1/stations \
>  -H 'cache-control: no-cache' \
>  -H 'content-type: application/json' \
>  -d '{"name":"Brooklyn","opening_time":"8:12:00",
> "closing_time":"18:23:00"}'

{"result":{"id":1,"name":"Brooklyn","opening_time":"8:12:00","closing_time":"18:23:00"}}
```

GET

```bash
CURL -X GET "http://localhost:8000/v1/stations/1"
```

DELETE

```bash
CURL -X DELETE "http://localhost:8000/v1/stations/1"
```

Gin log

```bash
[GIN] 2023/03/05 - 17:59:25 | 200 |   24.046913ms |             ::1 | POST     "/v1/stations"
[GIN] 2023/03/05 - 18:02:31 | 200 |      3.8567ms |             ::1 | GET      "/v1/stations/1"
[GIN] 2023/03/05 - 18:02:57 | 200 |      6.2644ms |             ::1 | DELETE   "/v1/stations/1"
```

### Revel

Revel 是一个很厚重的 web 框架，对比 django。因此简介是无法了解这个框架的。

install revel:

```bash
go get github.com/revel/revel
go get github.com/revel/cmd/revel
```

run

```bash
revel run github.com/narenaryan/railAPIRevel
```

test api:

```bash
CURL -X GET "http://127.0.0.1:8000/v1/trains/1"
```

总结：

- 当您需要端到端时使用 Revel.go Web 应用程序（模板和 UI），
- 使用 Gin 快速创建 REST 服务，
- 并在 API 的性能至关重要时使用 go‑rest

## chapter 05

install mongodb org in wsl2 ubuntu22.04:
> https://www.mongodb.com/docs/manual/tutorial/install-mongodb-on-ubuntu/

```bash
go get gopkg.in/mgo.v2
```

Boosting the querying performance with indexing

利用索引加速数据查询。

创建 index:

```bash
db.movies.createIndex({year: 1})
```

```bash
db.movies.find({year: {$lt: 2010}}).explain("executionStats")
```

---

迷你电商购物网站系统。

## chapter 06

Working with Protocol Buffers and GRPC

- Protocol buffers introduction
- Format of the protocol buffers
- Compilation process of a protobuf
- GRPC, a modern RPC library
- Bidirectional streaming with GRPC

install proto3.3:

```bash
# Make sure you grab the latest version
curl -OL https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip
# Unzip
unzip protoc-3.3.0-linux-x86_64.zip -d protoc3
# Move only protoc* to /usr/bin/
sudo mv protoc3/bin/protoc /usr/bin/protoc
```

compile .proto example file

```bash
protoc --go_out=. *.proto
```

会遇到缺失 protoc-gen-go 的问题，解决办法：

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# check your GOPATH, make sure have protoc-gen-go 
```

最后，关键一步，将你的 GOPATH 添加到系统变量（注意不是用户变量）。

然后使用 Chapter06/protobufs 例子代码进行测试。很容易生成 json 格式的字符串表示。

protobuf 比 json 的优势在于 protobuf 更加轻量，而且也能轻易地转化为 json 表示。

### GRPC

- JSON RPC => 专门传输 JSON 的 RPC 应用。
- Google RPC => 专门传输 protocol buffers 的 RPC 应用

```bash
go get google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

GRPC的优点：
- GRPC uses HTTP/2, which is a binary protocol => HTTP2
- Header compression is possible in HTTP/2, which means less overhead => 头部压缩
- We can multiplex many requests on one connection => 连接复用
- Usage of protobufs for strict typing of data => 强数据类型
- Streaming of requests or responses is possible instead of request/response transactions => 流化

```bash
#protoc -I datafiles/ datafiles/transaction.proto --go_out=plugins=grpc:datafiles

# not work because of compatibility: mustEmbedUnimplementedMoneyTransactionServer
protoc -I datafiles/ datafiles/transaction.proto --go-grpc_out=:datafiles
```
书中这行命令无法执行。解决方案有两个：

第一个：在编译时指定不强制implement servers，放宽限制。
```bash
# work
protoc -I datafiles/ datafiles/transaction.proto --go-grpc_out=require_unimplemented_servers=false:datafiles
```

第二个：在server side明确指示方法不实现。
```
// server is used to create MoneyTransactionServer.
type server struct {
	pb.UnimplementedMoneyTransactionServer
}
```
关联源码：
```
// UnimplementedMoneyTransactionServer must be embedded to have forward compatible implementations.
type UnimplementedMoneyTransactionServer struct {
}
```

详细代码参考：Chapter06/grpc_example

---

双向消息交流。
