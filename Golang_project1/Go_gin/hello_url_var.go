package Go_gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloUrlVar() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from v1!")
		})
		v1.GET("/goodbye", func(c *gin.Context) {
			c.String(http.StatusOK, "Goodbye from v1!")
		})

		v1.GET("/hello/:name", func(c *gin.Context) { //通过变量传递参数
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		v1.GET("/goodbye/:name/:action", func(c *gin.Context) { //通过变量传递多个参数
			name := c.Param("name")
			action := c.Param("action")
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Goodbye %s, action: %s", name, action),
			})
		})

		v1.POST("/form_post/:name/action", func(c *gin.Context) {
			message := c.PostForm("message")
			name := c.Param("name")
			nick := c.DefaultPostForm("nick", "anonymous")

			c.JSON(http.StatusOK, gin.H{
				"status":  "posted",
				"message": message,
				"nick":    nick,
				"name":    name,
			})
		})

		v1.DELETE("/delete/:name/*action", func(c *gin.Context) { //*action表示匹配后续所有的路径
			name := c.Param("name")
			action := c.Param("action")
			c.JSON(http.StatusOK, gin.H{
				"name":   name,
				"action": action,
				"status": "deleted",
			})
			c.String(http.StatusOK, "Delete action performed")
		})
	}

	r.Run(":8080")
}
