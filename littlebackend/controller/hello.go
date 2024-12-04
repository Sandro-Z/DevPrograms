package controller

import (
	"fmt"
	"net/http"
	"time"

	"exam_go/service"

	"github.com/gin-gonic/gin"
)

type HelloController struct {
}

func (s HelloController) Hello(c *gin.Context) {

	helloService := service.HelloService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    helloService.Hello(),
		Message: "welcome~",
	})
}

func (s HelloController) HelloTime(c *gin.Context) {
	var form struct {
		Date time.Time `form:"date" binding:"required,timing" time_format:"2006-01-02"`
	}
	if err := c.ShouldBindQuery(&form); err != nil {
		fmt.Printf("controller %v", err)
		c.Error(&gin.Error{
			Err:  err,
			Type: service.ParamErr,
		})
		return
	}

	helloService := service.HelloService{}

	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    helloService.HelloTime(form.Date),
	})
}
