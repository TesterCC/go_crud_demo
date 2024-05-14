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
go build
./go_crud_demo
```

cross compiling
```shell
// 交叉编译+减少文件大小, linux amd64
GOOS="linux" GOARCH="amd64" go build -o go_crud_demo -ldflags "-w -s" main.go

```

完整项目编译可执行文件，以常见的三种系统为例：
```
1. Linux
GOOS="linux" GOARCH="amd64" go build -o go_crud_demo -ldflags "-w -s"

2. Windows
GOOS="windows" GOARCH="amd64" go build -o go_crud_demo.exe -ldflags "-w -s"

3. MacOS(Intel CPU)
GOOS="darwin" GOARCH="amd64" go build -o go_crud_demo -ldflags "-w -s"

4. MacOS(ARM CPU)
GOOS="darwin" GOARCH="arm64" go build -o go_crud_demo -ldflags "-w -s"
```

## Deployment

- 后端可执行文件 go_crud_demo
- static目录中的静态资源数据


