# cobra
A simple and usable wild Jane assembled by the cobra frame and viper configuration

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