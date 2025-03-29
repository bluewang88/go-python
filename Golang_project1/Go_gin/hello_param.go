package Go_gin

//从get和post获取参数

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloGetParam() {
	r := gin.Default()

	// This handler will match /user?name=xxx&message=xxx
	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Guest") // 如果没有传递name参数，则使用Guest作为默认值
		message := c.Query("message")           // 是 c.Request.URL.Query().Get("lastname") 的简写
		c.JSON(http.StatusOK, gin.H{
			"Hello":   name,
			"Message": message,
		})
	})

	//使用PostForm从表单中获取参数，body中的form-data参数
	r.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//混合使用get和post获取参数
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		c.JSON(http.StatusOK, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	r.Run(":8080")
}
