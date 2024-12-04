package service

import (
	"errors"
	"exam_go/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AdminService struct {
}

type AdminRestsResponse struct {
	Total int64            `json:"total"`
	List  []AdminRestsList `json:"list"`
}

type AdminRestsList struct {
	ID       int    `json:"ID,omitempty"`
	Name     string `json:"restaurantname,omitempty"`
	Status   string `json:"status,omitempty"`
	Describe string `json:"describe,omitempty"`
	Userid   string `json:"userid,omitempty"`
}

func (h AdminService) AdminHelloService() string {
	return fmt.Sprintln("你是管理员哟,请访问其他接口来操作吧")
}

func (a AdminService) HandleGetNewRestListService(c *gin.Context) (resp *AdminRestsResponse, err error) {
	var uri struct {
		Page  int `form:"page" json:"page" uri:"page" binding:"required" validate:"gte:1"`
		Limit int `form:"limit" json:"limit" uri:"limit" binding:"required"`
	}

	if err := c.ShouldBindQuery(&uri); err != nil {
		return nil, err
	}
	resp = &AdminRestsResponse{}
	println(uri.Page)
	println(uri.Limit)
	offest := (uri.Page - 1) * uri.Limit
	adminrestsList := &[]AdminRestsList{}
	var total int64
	tx := model.DB.Model(&model.Restaurant{})
	tx.Limit(uri.Limit).Offset(offest)
	if err := tx.Where("status = ?", 0).Find(adminrestsList).Count(&total).Error; err != nil {
		return nil, err
	}
	resp.Total = total
	resp.List = *adminrestsList
	return resp, nil
}

func (a AdminService) HandleAdminGetRestFoodListService(c *gin.Context) (resp *SearchFoodsResponse, err error) {
	var uri struct {
		Page  int `form:"page" json:"page" uri:"page" binding:"required" validate:"gte:1"`
		Limit int `form:"limit" json:"limit" uri:"limit" binding:"required"`
		ID    int `form:"id" json:"id" uri:"id" binding:"required"`
	}

	if err := c.ShouldBindQuery(&uri); err != nil {
		return nil, err
	}
	resp = &SearchFoodsResponse{}
	println(uri.Page)
	println(uri.Limit)
	offest := (uri.Page - 1) * uri.Limit
	restfoodList := &[]FoodsList{}
	var total int64
	tx := model.DB.Model(&model.Food{})
	tx.Limit(uri.Limit).Offset(offest)
	if err := tx.Where("resid = ?", uri.ID).Find(restfoodList).Count(&total).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	resp.Total = total
	resp.List = *restfoodList
	if len(resp.List) == 0 {
		return nil, errors.New("没有找到该商家")
	}
	return resp, nil
}

func (a AdminService) HandleChange(c *gin.Context) (err error) {
	var uri struct {
		Status int `form:"status" json:"status" uri:"limit" binding:"required"`
		ID     int `form:"id" json:"id" uri:"id" binding:"required"`
	}

	if err := c.ShouldBind(&uri); err != nil {
		return err
	}
	tx := model.DB.Model(&model.Restaurant{})
	if err := tx.Where("id = ?", uri.ID).Update("status", uri.Status).Error; err != nil {
		return err
	}

	return nil
}
