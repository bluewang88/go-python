package Go_gin

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 中间件

func Mylogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在请求之前执行的逻辑
		// 例如，记录请求信息
		t := time.Now()
		c.Writer.Header().Set("X-Custom-Header", "This is a custom header")

		log.Printf("Request time: %v", t)
		c.Next() // 调用下一个中间件或处理函数
	}
}

func HelloMiddleware() {
	// r := gin.Default() // 创建一个默认的路由器, 默认使用Logger和Recovery(crash-free)中间件,WSGI（Web Server Gateway Interface）

	r := gin.New()        // 创建一个新的路由器, 不使用默认的中间件
	r.Use(gin.Logger())   // 使用Logger中间件
	r.Use(gin.Recovery()) // 使用Recovery中间件
	// r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
	// 	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" \"%s\"\n",
	// 		param.ClientIP,
	// 		param.TimeStamp.Format("2006-01-02 15:04:05"),
	// 		param.Method,
	// 		param.Path,
	// 		param.Request.Proto,
	// 		param.StatusCode,
	// 		param.Latency,
	// 		param.Request.UserAgent(),
	// 		param.ErrorMessage,
	// 	)
	// }))
	// r.Use(gin.RecoveryWithWriter(os.Stdout)) // 使用Recovery中间件,将错误信息输出到标准输出

	// authrized := r.Group("/auth")

	// 定义一个中间件
	r.Use(func(c *gin.Context) {
		// 在请求之前执行的逻辑
		c.Writer.Header().Set("X-Custom-Header", "This is a custom header")
		c.Next() // 调用下一个中间件或处理函数
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // 运行在本地服务器上的8080端口
}
