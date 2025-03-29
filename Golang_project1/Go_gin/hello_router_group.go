package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloRouteGroup() {
	r := gin.Default()

	//Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from v1!")
		})
		v1.GET("/goodbye", func(c *gin.Context) {
			c.String(http.StatusOK, "Goodbye from v1!")
		})
	}

	//Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from v2!")
		})
		v2.GET("/goodbye", func(c *gin.Context) {
			c.String(http.StatusOK, "Goodbye from v2!")
		})
	}

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}
}
