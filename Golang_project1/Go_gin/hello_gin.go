package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloGin() {
	r := gin.Default() // 创建一个默认的路由器, 默认使用Logger和Recovery中间件,WSGI（Web Server Gateway Interface）

	r.GET("/ping", func(c *gin.Context) { // 定义路由
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // 运行在本地服务器上的8080端口
}
