## Environment

### Backend
go + gin + mysql5.7/8.0 

### Frontend

```
local debug:
~/ws_vuejs/vue3-demo

npm run dev
```

## Build Binary

```
go build -o go_crud_demo
./v
```

cross compiling
```shell
// 交叉编译+减少文件大小, linux amd64
GOOS="linux" GOARCH="amd64" go build -o go_crud_demo -ldflags "-w -s" main.go
```

## Deployment

- 后端可执行文件 go_crud_demo
- static目录中的静态资源数据


