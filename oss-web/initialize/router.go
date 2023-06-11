package initialize

import (
	"go-shop-api/oss-web/middlewares"
	"go-shop-api/oss-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	Router.LoadHTMLFiles("oss-web/templates/index.html")
	// 配置静态文件夹路径 第一个参数是api，第二个是文件夹路径
	Router.StaticFS("/static", http.Dir("oss-web/static"))
	Router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "posts/index",
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/oss/v1")
	router.InitOssRouter(ApiGroup)

	return Router
}
