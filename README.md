# cobra
A simple and usable wild Jane assembled by the cobra frame and viper configuration

## 项目结构：
--cmd corba cli命令行服务  
--configs 配置  
--router  http路由管理
--pkg 
-- | -- db    mysql定义
-- | -- redis redis定义

## 程序说明:
>  `go run main.go server --config=.cobra.yaml` 命令行形式启动对应服务 `server` 表示启动`cmd`目录下对应的该服务