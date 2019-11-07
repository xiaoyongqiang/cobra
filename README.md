# cobra

1. 由 `cobra` 应用程序，生成的应用框架, 可通过 `cobra add {app}` 生成需要命令行服务  开源地址 `github.com/spf13/cobra`

2. 框架中引用 `github.com/spf13/viper` 加载配置包，支持yaml, json... 并且实时监听实现服务热更新

3. `cmd` 目录下，对应的命令行服务处理不同的业务程序。 比如 `server` 服务 程序中是一个运用 `gin` 为列子的http服务

4. `cmd` 目录下， `nsq` 消息队列服务的使用， 创建topic、channel 监听集群管理的nsq服务的消息， 处理异步消息，本地安装nsq服务可通过 `godep get github.com/bitly/nsq/...` 安装编译包

5. `cmd` 目录下， `grpc` 框架的使用:
    
    * 先进行安装 `go get github.com/grpc/grpc-go`
    * 其次我们要下载 `protocbuf生成器`
    * 然后安装通过 `go get github.com/golang/protobuf`
    * 接着我们通过 `go get -u github.com/golang/protobuf/protoc-gen-go` 生成可执行文件 `protoc-gen-go`（ 将下载 `protocbuf生成器` 文件中 `protoc.exe` 设置为全局变量，将生成的 `protoc-gen-go.exe` 移动在该文件bin下面，全局使用得将设为全局变量 `PATH=$PATH:xxx/protoc/bin` ）
    * 接着就可以书写你需要的 `proto` 文件，通过指令 `protoc -I .  gateway.proto --go_out=plugins=grpc:.` 最后在当前文件夹生成pb.go文件...

## 项目结构：

--cmd     程序入口  
--router  路由管理  
--configs 配置  
--tools   工具包  
--nsqpubsub 消息队列处理  
--apis   
-- | -- handler 业务处理  
-- | -- models 数据库模型定义  
--pkg   
-- | -- db   
-- | -- redis   
-- | -- nsq   
  

### 程序启动:

 > `go run main.go server --config=.cobra.yaml` 命令行形式启动 `第一参数` 为对应服务 `server` 表示启动 `cmd` 目录下对应的该服务 `arg[]` 命令行参数根据需求可配


### Api 规范

1. 路由规则
> `/大版本号/模块/接口?format=响应格式 ` ,大版本号当前为`v1`;format可忽略，默认为json,目前只实现json  

2. 接口说明  
除登录接口外，所有请求头部必须包含以下共同参数：

    * `X-UID` 用户id,由登录接口返回
    * `X-SIGN` 签名
    * `X-TS` 请求时间戳

3. 接口签名
> 采用 `md5` 加签验证, 协议头参数按照 `{X-TS}{X-UID}{token}` 拼接进行加签， 得到的数据作为协议 `X-SIGN` 的参数值


4. 返回值说明
若无特殊说明，接口返回值均采用以下形式的json返回 
```json
  {
     "status" : 0, 
     "msg" : "ok", 
     ...
  }
```
其中  
状态码`status` 为0，表示成功，非0表示失败；失败描述查看`msg`值