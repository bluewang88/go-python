package Go_gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//表单验证,login

// 表单格式请求
// curl --location 'localhost:8080/loginjson' \
// --form 'nameform="wangwe"' \
// --form 'passwordform="123456"'

// JSON格式请求
// curl --location 'localhost:8080/loginjson' \
// --header 'Content-Type: application/json' \
// --data '{
//     "namejson": "wangwei",
//     "passwordjson": "123456"
// }'

// XML格式请求
// curl --location 'localhost:8080/loginjson' \
// --header 'Content-Type: application/xml' \
// --data '<root>
//     <namexml>wangwei</namexml>
//     <passwordxml>123456</passwordxml>
// </root>'

type LoginForm struct {
	Name     string `form:"nameform" json:"namejson" xml:"namexml" binding:"required,min=5,max=10"`
	Password string `form:"passwordform" json:"passwordjson" xml:"passwordxml" binding:"required"`
}

func HelloFormValidationLogin() {
	r := gin.Default()
	r.POST("/loginjson", func(c *gin.Context) {
		var loginform LoginForm

		if err := c.ShouldBind(&loginform); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})
	r.Run(":8080")
}
