package routes

import (
	_ "gebi/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func routeWithoutMiddleWare() [][3]interface{} {
	//请求类型，请求url，请求方法
	return [][3]interface{}{
		{http.MethodGet, "swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)},
		{http.MethodGet, "address/list", ctrl.address.List},
		{http.MethodPost, "address/add", ctrl.address.Add},
		{http.MethodPost, "address/update", ctrl.address.Update},
		{http.MethodPost, "address/del", ctrl.address.Del},
		{http.MethodPost, "user/register", ctrl.user.Register},
		{http.MethodPost, "user/login", ctrl.user.Login},
		{http.MethodGet, "user/wechat/login", ctrl.wechat.WechatLogin},
		{http.MethodPost, "ft/changeagepic", ctrl.ft.ChangeAgePic},
	}
}
