package core

import (
	"bytes"
	"context"
	"fmt"
	"gebi/app/Http/Middleware"
	"gebi/app/Http/Serializer"
	"gebi/routes"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

var Router *gin.Engine

// SetupRouter 路由
func SetupRouter() *gin.Engine {
	Router = gin.New()
	Router.Use(logFormatter())
	Router.Use(Middleware.PageSize())
	Router.Use(customRecover())
	route := Router.Group("")
	routing(route, routes.ApiRouters())
	authRouting(route, routes.AuthRouters())
	pprof.Register(Router) //性能分析
	return Router
}

func initRoute() {
	gin.DefaultWriter = io.MultiWriter(ginLogFile())
	SetupRouter()
	shutdownGraceFull()
}

func shutdownGraceFull() {
	srv := &http.Server{
		Addr:    ":8081",
		Handler: Router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	s := <-quit
	log.Println("Shutting down server...", s)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

//系统请求日志文件
func ginLogFile() *os.File {
	f, _ := os.OpenFile("gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return f
}

//格式化日志
func logFormatter() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

//权限路由
func authRouting(router *gin.RouterGroup, middleRoutes []routes.AuthMiddleWareRoute) {
	r := make([]*gin.RouterGroup, len(middleRoutes))
	for index, item := range middleRoutes {
		r[index] = router.Group("")
		for _, authorized := range item.MiddleWares {
			r[index].Use(authorized)
		}
		routing(r[index], item.Uris)
	}
}

//用户自定义recover
func customRecover() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(Serializer.Response); ok {
			log.Errorf("SYSTEM ACTION PANIC: %v, stack: %v", err, string(debug.Stack()))
			c.JSON(http.StatusInternalServerError, err)
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}

func uriString(url string) string {
	uri := bytes.Buffer{}
	prefix := os.Getenv("ROUTE_PREFIX")
	if prefix != "" {
		uri.WriteString(prefix)
		uri.WriteString("/")
	}
	uri.WriteString(url)

	return uri.String()
}

//路由派发
func routing(router *gin.RouterGroup, routes [][3]interface{}) {
	for _, item := range routes {
		uri := uriString(item[1].(string))
		handler := getRouteTypeFunc(item[2])
		switch item[0] {
		case http.MethodGet:
			router.GET(uri, handler)
		case http.MethodDelete:
			router.DELETE(uri, handler)
		case http.MethodPost:
			router.POST(uri, handler)
		case http.MethodPatch:
			router.PATCH(uri, handler)
		case http.MethodPut:
			router.PUT(uri, handler)
		case http.MethodOptions:
			router.OPTIONS(uri, handler)
		case http.MethodHead:
			router.HEAD(uri, handler)
		case "Any":
			router.Any(uri, handler)
		default:
			log.Info("请求方法错误！，请重新定义")
		}
	}
}

func getRouteTypeFunc(item interface{}) gin.HandlerFunc {
	if value, ok := item.(gin.HandlerFunc); ok {
		return value
	} else {
		return item.(func(ctx *gin.Context))
	}
}
