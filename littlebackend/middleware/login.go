package middleware

import (
	"errors"
	"exam_go/model"
	"exam_go/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") == nil {
		c.Error(&gin.Error{
			Err:  errors.New("未登录"),
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	c.Next()
}

func IsNotLogin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") != nil {
		c.Error(&gin.Error{
			Err:  errors.New("请勿重复登陆,如需登录其他账号请先登出"),
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	c.Next()
}

func IsRest(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") == nil {
		c.Error(&gin.Error{
			Err:  errors.New("未登录"),
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	rest := &model.Restaurant{}
	if err := model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		c.Error(&gin.Error{
			Err:  err,
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	c.Next()
}

func IsAdmin(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userId") == nil {
		c.Error(&gin.Error{
			Err:  errors.New("未登录"),
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	rest := &model.User{}
	if err := model.DB.Where("id = ? AND status = ?", session.Get("userId").(int), 1).First(&rest).Error; err != nil {
		c.Error(&gin.Error{
			Err:  err,
			Type: service.AuthErr,
		})
		c.Abort()
		return
	}
	c.Next()
}
