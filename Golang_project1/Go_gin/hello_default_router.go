package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloGinDefaultRouter() {
	r := gin.Default() // 创建一个默认的路由器, 默认使用Logger和Recovery(crash-free)中间件,WSGI（Web Server Gateway Interface）

	r2 := gin.New() // 创建一个没有默认中间件的路由器，没有日志和恢复中间件

	r.GET("/ping", func(c *gin.Context) { // 定义路由
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r2.GET("/ping", func(c *gin.Context) { // 定义路由
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // 运行在本地服务器上的8080端口

	r2.Run(":8081") // 运行在本地服务器上的8081端口
}
