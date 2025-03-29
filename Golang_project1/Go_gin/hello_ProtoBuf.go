package Go_gin

// 返回protobuf格式的数据

import (
	"Golang_project1/Go_gin/proto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloProtoBuf() {

	r := gin.Default()

	r.GET("/protobuf", func(c *gin.Context) {

		user := proto.Teacher{
			Name:    "Rwxsht",
			Age:     18,
			Subject: "math",
			Email:   "root@163.com",
		}
		c.ProtoBuf(http.StatusOK, &user)
	})

	r.Run(":8080")
}
