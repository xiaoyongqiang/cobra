# cobra
A simple and usable wild Jane assembled by the cobra frame and viper configuration

## 项目结构：
--cmd 命令行服务   
--configs 配置 
--router 路由管理 
--pkg   
-- | -- redis   
-- | -- db  数据库


## 项目结构：
--api 接口路由定义, 业务实现在internal/handler中  
--cmd 程序入口  
--configs 配置  
--internal   
-- | -- handler 业务处理  
-- | -- models 数据库模型定义  
--pkg 
-- | -- types 常用变量的定义  

## 程序说明:

 > `go run main.go server --config=.cobra.yaml` 命令行形式启动对应服务 `server` 表示启动`cmd`目录下对应的该服务