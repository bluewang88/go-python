package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//表单验证, 通过表单提交数据，后端程序验证前端传递过来的表单数据是否合法

func HelloFormValidation(c *gin.Context) {
	// 获取表单内容
	name := c.PostForm("name")
	age := c.PostForm("age")

	// 验证表单内容
	if name == "" {
		c.String(http.StatusBadRequest, "name is required")
		return
	}

	if age == "" {
		c.String(http.StatusBadRequest, "age is required")
		return
	}

	// 验证通过，返回结果
	c.String(http.StatusOK, "Hello, %s, you are %s years old", name, age)
}
