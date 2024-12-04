package controller

import (
	"fmt"
	"net/http"

	"exam_go/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (s UserController) HandleUserinfo(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.Userinfo(c)

	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    resp,
	})
}

func (u UserController) HandleSignup(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleSignupservice(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleLogin(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleLoginService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleLogout(c *gin.Context) {
	session := sessions.Default(c)
	var resp struct {
		Login bool `json:"login"`
	}
	if session.Get("userId") == nil {
		println(session.Get("userId"))
		resp.Login = false
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    resp,
		})
		return
	}
	session.Clear()
	if err := session.Save(); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	resp.Login = true
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleGetRestList(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleGetRestListService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleGetRestFoodList(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleGetRestFoodListService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleGetFoodInfo(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleGetFoodInfoService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleOrderFood(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleOrderFoodService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleOrderInfo(c *gin.Context) {
	userService := service.UserService{}
	resp, err := userService.HandleGetOrderInfoService(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    resp,
	})
}

func (u UserController) HandleLikeInfo(c *gin.Context) {
	userService := service.UserService{}
	err := userService.HandleLikeFood(c)
	if err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.SysErr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
