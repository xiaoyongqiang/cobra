# cobra

1. 由 `cobra` 应用程序，生成的应用框架, 可通过 `cobra add {app}` 生成需要命令行服务  开源地址 `github.com/spf13/cobra`

2. 框架中引用 `github.com/spf13/viper` 加载配置包，支持yaml, json... 并且实时监听实现服务热更新

3. `cmd` 目录下，对应的命令行服务处理不同的业务程序。 比如 `server` 服务 程序中是一个运用 `gin` 为列子的http服务

## 项目结构：

--cmd     程序入口  
--router  路由管理  
--configs 配置  
--tools   工具包  
--apis   
-- | -- handler 业务处理  
-- | -- models 数据库模型定义  
--pkg   
-- | -- db   
-- | -- redis   
  

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