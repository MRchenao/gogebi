package routes

import (
	"net/http"
)

func wechatMiddleWareRoutes() [][3]interface{} {
	return [][3]interface{}{
		{http.MethodGet, "user/wechat/me", ctrl.wechat.WechatMe},
	}
}
