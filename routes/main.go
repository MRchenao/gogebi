package routes

import (
	"gebi/app/Http/Controllers/Api"
	"gebi/app/Http/Middleware"
	"github.com/gin-gonic/gin"
)

//控制器类
var ctrl struct {
	address Api.Address
	user    Api.User
	wechat  Api.Wechat
	ft  Api.Ft
}

type AuthMiddleWareRoute struct {
	MiddleWares []gin.HandlerFunc
	Uris        [][3]interface{}
}

//无需验证的路由
func ApiRouters() [][3]interface{} {
	return routeWithoutMiddleWare()
}

//需要中间件验证的路由
func AuthRouters() []AuthMiddleWareRoute {
	return []AuthMiddleWareRoute{
		{MiddleWares: []gin.HandlerFunc{Middleware.Auth()}, Uris: authMiddleWareRoutes()},
		{MiddleWares: []gin.HandlerFunc{Middleware.Wechat()}, Uris: wechatMiddleWareRoutes()},
	}
}
