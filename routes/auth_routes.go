package routes

import "net/http"

func authMiddleWareRoutes() [][3]interface{} {
	return [][3]interface{}{
		{http.MethodGet, "user/me", ctrl.user.UserMe},
	}
}
