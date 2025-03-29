package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloJson() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/morejson", func(c *gin.Context) {

		var msg struct {
			Name    string `json:"user"` // 通过json:"user"定义json的key为user，没有惠子个则json的key为Name
			Message string
			Number  int
		}

		msg.Name = "root"
		msg.Message = "测试json"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
