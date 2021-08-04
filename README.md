# tudo-be

## 工程结构
+ .github

  github actions相关配置文件

+ config

  服务端相关配置文件

+ controller

  http请求的初步处理

+ service

  具体业务的处理

+ model

  各类模型

+ model/dao

  数据库及缓存相关模型(数据操作对象)

+ model/dto

  客户端请求相关模型(数据传输对象)

+ util

  通用工具

+ docs

  API文档生成相关配置，由swag init命令自动生成

## API文档

访问/swagger/index.html

## 主要使用的技术/工具/框架
+ git
+ docker、docker-compose
+ MySQL、redis
+ Nginx
+ gin、gorm、go-redis、gin-swagger