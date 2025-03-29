package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloGinHttpFunc() {
	r := gin.Default()

	// 定义路由
	r.GET("/helloget", func(c *gin.Context) {
		c.String(http.StatusOK, " GET Hello World!")
	})
	r.POST("/hellopost", func(c *gin.Context) {
		c.String(http.StatusOK, " POST Hello World!")
	})
	r.PUT("/helloput", func(c *gin.Context) {
		c.String(http.StatusOK, " PUT Hello World!")
	})
	r.DELETE("/hellodelete", func(c *gin.Context) {
		c.String(http.StatusOK, " DELETE Hello World!")
	})
	r.PATCH("/hellopatch", func(c *gin.Context) {
		c.String(http.StatusOK, " PATCH Hello World!")
	})
	r.HEAD("/hellohead", func(c *gin.Context) {
		c.String(http.StatusOK, " HEAD Hello World!")
	})
	r.OPTIONS("/hellooptions", func(c *gin.Context) {
		c.String(http.StatusOK, " OPTIONS Hello World!")
	})

	r.Run(":8080")
}
