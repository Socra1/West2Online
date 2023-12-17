package midddleware

import (
	"net/http"
	"time"
	utils "todo_list/token"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			c.JSON(400, gin.H{
				"status": code,

				"data": "缺少Token",
			})
			c.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = 404
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 403
		}

		if code != 200 {
			if code == 404 {
				c.JSON(400, gin.H{
					"status": code,
					"data":   "未知错误",
				})
			} else if code == 403 {
				c.JSON(400, gin.H{
					"status": code,
					"data":   "身份过期，请重新登录",
				})

			}
			c.Abort()
			return
		}
		c.Next()
	}
}
