package controller

import (
	"fmt"
	"net/http"

	"exam_go/service"

	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type RestController struct {
}

func (s RestController) HandleRestinfo(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.Restinfo(c)

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

func (s RestController) HandleRestFoodAdd(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.HandleFoodAddservice(c)

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

func (s RestController) HandleNewRest(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.HandleNewRestservice(c)

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

func (s RestController) HandleNewRestinfo(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.HandleNewRestinfoService(c)

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

func (s RestController) HandleCheck(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.HandleCheckService(c)

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

func (s RestController) HandleChange(c *gin.Context) {
	restService := service.RestService{}
	err := restService.HandleChangeService(c)

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
	})
}

func (s RestController) HandleChangeFood(c *gin.Context) {
	restService := service.RestService{}
	resp, err := restService.HandleChangeFoodService(c)

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
