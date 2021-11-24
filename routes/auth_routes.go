package routes

import "net/http"

func authMiddleWareRoutes() []Routes {
	return []Routes{
		{http.MethodGet, "user/me", ctrl.user.UserMe},
	}
}
