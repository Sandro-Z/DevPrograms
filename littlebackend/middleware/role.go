package middleware

import (
	"errors"

	"exam_go/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CheckRole(min int) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user-id")
		if id == nil {
			c.Error(&gin.Error{
				Err:  errors.New("未登录"),
				Type: service.AuthErr,
			})
			c.Abort()
			return
		}
		if id.(int) < min {
			c.Error(&gin.Error{
				Err:  errors.New("权限不足"),
				Type: service.LevelErr,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
