package Middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
)

var Page int
var Limit int

func PageSize() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Page, _ = strconv.Atoi(c.DefaultQuery("page", "1")); Page < 1 {
			Page = 1
		}

		if Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "10")); Limit < 1 || Limit > 100 {
			Limit = 10
		}

		log.Info("page:", Page, Limit)
		c.Next()
	}
}
