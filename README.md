## 这是什么?
>   1.这是一个基于go语言gin框架的web项目骨架，专注于前后端分离的业务场景,其目的主要在于将web项目主线逻辑梳理清晰，最基础的东西封装完善，开发者更多关注属于自己的的业务即可。  
>   2.本项目骨架封装（主要包括用户相关的接口参数验证器、注册、登录获取token、自动化配置、优雅的路由风格、简易的数据库操作CURD以及JWT鉴权、单元测试、grpc、swagger文档等），开发者拉取本项目骨架，在此基础上就可以快速开发自己的项目。  
>   3.后续待完善功能为后台管理权限，角色用户等功能的开发，框架优化等

## 项目目录结构介绍
```jsunicoderegexp
├── app                                项目应用目录
│   ├── Admin                          后台管理模块
│   │   ├── Controllers                后台管理控制器  
│   │   ├── Middleware                 后台管理中间件（存放鉴权，中间事务处理等逻辑文件）
│   │   └── Services                   后台业务逻辑服务模块
│   ├── Http                           前端api模块
│   │   ├── Controllers          
│   │   │   └── Api                    api模块控制器存放文件位置，以下几个为样例文件，实际开发中可根据样例开发，不用可删除
│   │   │       ├── address.go           
│   │   │       ├── ft.go
│   │   │       ├── main.go            api控制器通用方法文件
│   │   │       └── user.go
│   │   ├── Middleware                  中间件模块
│   │   │   ├── auth.go                 用户权限验证
│   │   │   └── pageSize.go             分页中间件       
│   │   ├── Serializer                 序列化输出模块
│   │   │   ├── address.go
│   │   │   ├── ft.go
│   │   │   ├── main.go          序列化输出公用方法文件
│   │   │   └──user.go
│   │   └── Services                   业务逻辑模块
│   │       ├── address.go
│   │       ├── ft.go
│   │       ├── main.go                业务逻辑公用方法文件
│   │       └── user.go
│   ├── Models                        数据库模型
│   │   ├── address.go
│   │   ├── Build                     数据库公用操作方法文件
│   │   │   └── build.go
│   │   └──user.go
│   └── Repositories                  数据仓库模块（获取数据，数据缓存等操作逻辑在这一层完成）
│       ├── address.go
│       └──user.go
├── config                                  业务配置文件
│   └── config.yml
├── core                                    框架核心逻辑模块
│   ├── bootstrap.go                  核心启动文件 
│   ├── log.go                        日志文件
│   ├── router.go                     核心路由处理文件
│   ├── rpc_service.go                rpc服务  
│   └── viper.go                      配置加载初始化  
├── docs                              swagger文档文件
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── gin.log                           项目日志文件
├── main.go                           项目入口文件
├── Makefile
├── proto                             proto文件
│   ├── hello.pb.go
│   ├── hello.proto
├── routes                             路由文件
│   ├── auth_routes.go           带权限路由
│   ├── main.go                  路由公用配置主文件
│   └── routes.go               不带权限路由
├── rpc
│   └── rpc_server.go           rpc服务模块
├── storage                           日志文件存储模块
│   └── logs
│       └── log.log
├── test                              单元测试模块
│   ├── address_test.go
│   └── main.go                 单元测试公用方法文件
└── utils                       工具模块目录
    ├── database                数据库操作
    │   └── db.go
    ├── redis_factory           redis操作
    │   └── redis.go
    ├── sign                    jwt加密鉴权
    │   └── jwt.go
    └──  util.go                 工具方法

```
##项目开发流程
### 基本流程
>   1.首先添加一个前端请求url，在routes/routes.go文件中添加您的url，格式如下：
各参数对应的含义分别是：请求类型，请求url，请求的控制器方法
```code
{http.MethodPost, "user/login", ctrl.user.Login}
```
>   2.在app/Http/Controllers/Api目录下创建控制器文件address.go添加login方法，具体请看项目文件

>   3.在app/Http/Services中创建对应的services文件，这里书写业务逻辑，同理在repositories创建操作数据的文件，在model中创建数据库文件

>   4.在app/Http/Serializer中创建最终输出的数据序列化处理逻辑文件，最终输出给前端，请求结束
### 针对需要鉴权的路由的处理
基本流程中展示的是无需鉴权操作的流程，针对需要鉴权的路由，首先您需要在routes/main.go文件的AuthRouters方法里面添加您的鉴权中间件，例如：
```code
{MiddleWares: []gin.HandlerFunc{Middleware.Auth()}, Uris: authMiddleWareRoutes()},
```
这里的MiddleWares就是需要添加的鉴权的中间件，如果有多个可以添加多个中间件。Uris定义了你需要通过这个中间件的路由函数方法，接下来就是在路由函数方法中添加您的路由，您的所有在函数方法中添加的路由都会首先需要先通过您的中间件，具体样例可参考auth_routes.go文件
### 数据库操作
本项目是居于gorm来操作数据库的，其中在app/Models/Build下封装了针对不定条件的查询公用方法：
```code
查询list
Build.BuildQueryList(wheres, []string{"*"}, "id desc").Find(&list)

构建where条件
Build.BuildWhere(database.DB, wheres)

不定条件更新
Build.BuildUpdates(addressModel, wheres, data)
```
其中where条件的构建格式如下：

```code
// BuildWhere 构建where条件
//1、and 条件 where := []interface{}{
//    []interface{}{"id", "=", 1},
//    []interface{}{"username", "chen"},
//}
//2、结构体条件  where := user.User{ID: 1, UserName: "chen"}
//3、in,or 条件 where := []interface{}{
//    []interface{}{"id", "in", []int{1, 2}},
//    []interface{}{"username", "=", "chen", "or"},
//}
//4、map条件  where := map[string]interface{}{"id": 1, "username": "chen"}
//5、and or混合条件  where := []interface{}{
//    []interface{}{"id", "in", []int{1, 2}},
//    []interface{}{"username = ? or nickname = ?", "chen", "yond"},
//}
```
##项目配置
>配置分两块，其中根目录下的.env文件主要配置一些项目基础配置如: 数据库，redis，路由前缀等的配置。
>另一块配置在config目录下，使用viper配置，主要配置一些业务逻辑方面的配置

## 项目启动
如果安装了make环境的直接使用
```shell
    make run
```
没有安装make环境的使用
```shell
    go run main.go
```
## 项目编译
如果安装了make环境的直接使用
```shell
make build
```
没有安装make环境的使用
```shell
go build
```
## 项目线上部署
项目默认的端口是8081，将编译后的文件直接拉到线上，创建.env文件，配置config文件，doc文档，storage日志文件即可，部署目录结构如下
```shell
config/
drwxr-xr-x  2 root root     4096 Nov 23 18:12 docs/
-rwxr-xr-x  1 root root      232 Nov 23 18:08 .env*
-rwxr-xr-x  1 root root 43204377 Nov 24 11:08 gebi*
-rw-r--r--  1 root root    78227 Nov 24 16:47 gin.log
drwxrwxrwx  3 root root     4096 Nov 23 17:11 storage/
```
