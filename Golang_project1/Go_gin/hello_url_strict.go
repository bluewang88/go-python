package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `uri:"name" binding:"required"`
	ID   string `uri:"id" binding:"required,uuid"`
}

func HelloUrlStrict() {
	r := gin.Default()
	r.GET("/user/:name/:id", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindUri(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"HTTPCode":   http.StatusBadRequest,
				"badrequest": "报错",
				"msg":        err,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"name": user.Name,
			"id":   user.ID,
		})
	})
	r.Run(":8080")
}
