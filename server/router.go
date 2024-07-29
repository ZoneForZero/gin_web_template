package server

import (
	"fmt"
	"gin_web_template/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	setTemplate(r)
	// 路由
	//v1 := r.Group("/api/v1")
	//{
	//	v1.POST("ping", api.Ping)
	//
	//	// 用户登录
	//	v1.POST("user/register", api.UserRegister)
	//
	//	// 用户登录
	//	v1.POST("user/login", api.UserLogin)
	//
	//	// 需要登录保护的
	//	auth := v1.Group("")
	//	auth.Use(middleware.AuthRequired())
	//	{
	//		// User Routing
	//		auth.GET("user/me", api.UserMe)
	//		auth.DELETE("user/logout", api.UserLogout)
	//	}
	//}
	return r
}
func setTemplate(r *gin.Engine) {
	// 加载模板文件
	r.LoadHTMLGlob("static/template/*")
	r.Static("css", "static/css")
	r.Static("img", "static/img")
	r.Static("js", "static/js")
	r.Static("font", "static/font")
	r.Static("file", "static/file")

	r.Static("fonts", "static/css/fonts")
	r.Static("webfonts", "static/webfonts")
	r.Static("picture", "static/picture")
	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})

	})

	// 定义路由
	r.GET("/:name", func(c *gin.Context) {
		htmlName := c.Param("name")
		fmt.Println(htmlName, htmlName == "")
		if htmlName == "" {
			htmlName = "index.html"
		}
		c.HTML(200, htmlName, gin.H{
			// "Name": "Gin User",
			//注意！：gin.H{"title": "Hello Gin"}实际是map[string]interface{}，通过用{{.Name}}实际就是获取了map中name字段的值，.就是上下文中gin.H传过来的数据；
		})
	})
}
