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
}

type Routes struct {
	Method     string
	Url        string
	Controller gin.HandlerFunc
}

type AuthMiddleWareRoute struct {
	MiddleWares []gin.HandlerFunc
	Uris        *[]Routes
}

//无需验证的路由
func ApiRouters() *[]Routes {
	return routeWithoutMiddleWare()
}

//需要中间件验证的路由
func AuthRouters() *[]AuthMiddleWareRoute {
	return &[]AuthMiddleWareRoute{
		{MiddleWares: []gin.HandlerFunc{Middleware.Auth()}, Uris: authMiddleWareRoutes()},
	}
}
