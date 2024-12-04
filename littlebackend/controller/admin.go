package controller

import (
	// "fmt"
	"fmt"
	"net/http"

	"exam_go/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func (a AdminController) AdminHello(c *gin.Context) {

	helloService := service.AdminService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    helloService.AdminHelloService(),
		Message: "welcome~",
	})
}

func (a AdminController) AdminRestList(c *gin.Context) {
	adminService := service.AdminService{}
	resp, err := adminService.HandleGetNewRestListService(c)

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

func (a AdminController) AdminRestFoodList(c *gin.Context) {
	adminService := service.AdminService{}
	resp, err := adminService.HandleAdminGetRestFoodListService(c)

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

func (a AdminController) AdminChange(c *gin.Context) {
	adminService := service.AdminService{}
	err := adminService.HandleChange(c)

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
