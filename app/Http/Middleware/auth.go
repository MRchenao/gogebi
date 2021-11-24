package Middleware

import (
	"gebi/app/Http/Serializer"
	"gebi/utils/sign"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

var jwtService = sign.JwtService{}

type authHeader struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := jwtService.ParseData(getTokenString(c))
		log.Info("用户id:", uid)
		c.Set("user_id", uid)
		c.Next()
	}
}

func getTokenString(c *gin.Context) string {
	header := authHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		Serializer.Err(400, "未授权", err)
	}
	/*	authorization := c.Request.Header.Get("Authorization")
		if authorization == "" {
			Serializer.Err(400, "未授权", nil)
		}*/
	parts := strings.Split(header.Authorization, " ")
	if len(parts) < 2 {
		Serializer.Err(400, "授权串错误", nil)
	}
	return parts[1]
}
