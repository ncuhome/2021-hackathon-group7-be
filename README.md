# 2021-hackathon-backend

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

+ document

  各类文档（包括API文档）

## 技术选择

### 主要使用的技术/工具/框架
+ git、github、github actions
+ docker、docker-compose
+ MySQL、redis
+ Nginx
+ gin、gorm、go-redis

### 选择原因

+ 团队内部对这些技术比较熟悉

+ 更换技术、框架等成本高

+ 后续容易找到接手的人员

## 主要实现的功能
+ 登录、注册
+ 绑定邮箱
+ 忘记密码（通过邮箱更改密码）
+ 按IP限制接口访问频率
+ 用户认证
+ 认证的用户发布活动
+ 用户对活动进行评论
+ 查看活动、评论等相关信息
+ 上传图片

### 计划实现的功能
+ QQ第三方登录
+ 搜索

### 功能细节
+ https
+ token鉴权,并具有主动废除token的能力
+ 密码加盐哈希存储
+ 使用redis对适当的数据缓存
+ 容器化测试/部署
+ 利用github actions和docker做CI/CD，持续集成、持续交付、持续部署
+ 部分功能模块化，可复用性强
+ 代码分层，整体耦合度低，易扩展易维护
+ 保存错误日志方便Debug
+ Nginx反向代理
+ 阿里云oss对象存储