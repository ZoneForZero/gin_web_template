package routers

import (
	controller "gin_web_template/controller"
	"gin_web_template/middleware"

	"os"

	"github.com/gin-gonic/gin"
)

// 简单应用，默认一个路由
func AllRouter() *gin.Engine {
	// 创建路由
	router := gin.Default()
	// 加载中间件, 顺序不能改
	router.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	router.Use(middleware.Cors())
	router.Use(middleware.CurrentUser())
	// 路由分组，testRouter的路由前面默认为/testRouter/
	testRouter := router.Group("/testRouter")
	{
		testRouter.GET("ping", controller.Ping)
		// 用户注册
		testRouter.POST("user/register", controller.UserRegister)
		// 用户登录
		testRouter.POST("user/login", controller.UserLogin)
		// 需要登录保护的
		auth := testRouter.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", controller.UserMe)
			auth.DELETE("user/logout", controller.UserLogout)
		}
	}
	return router
}
