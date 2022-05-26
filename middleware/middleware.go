package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token") //user token
		fmt.Println("token:", token)
		c.Next()
	}
}

// 跨域问题解决  不需要用这个
func Cros() gin.HandlerFunc {
	return func(c *gin.Context) {
		// method := c.Request.Method
		// origin := c.Request.Header.Get("Origin") //请求头部
		// if origin != "" {
		// 	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		// 	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		// 	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token, session")
		// 	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
		// 	c.Header("Access-Control-Max-Age", "172800")
		// 	c.Header("Access-Control-Allow-Credentials", "true")
		// }

		// //允许类型校验
		// if method == "OPTIONS" {
		// 	c.JSON(200, "ok!")
		// }

		c.Next()
	}
}
