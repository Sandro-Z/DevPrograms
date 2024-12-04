package service

import (
	// "errors"
	"errors"
	"exam_go/model"
	"fmt"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "gorm.io/gorm"
)

type RestService struct {
}

type Restinforesp struct {
	UserId int    `json:"userId"`
	Number string `json:"number"`
	Status int    `json:"status"`
}

type Foodinforesp struct {
	ID       int    `json:"id"`
	Foodname string `json:"foodname"`
	Describe string `json:"describe"`
	Cost     int    `json:"cost"`
}

type OrderInfo struct {
	ID       int    `json:"id"`
	Foodname string `json:"foodname"`
	Resid    int    `json:"resid"`
	Userid   int    `json:"userid"`
}

type Orderresp struct {
	Total int64       `json:"total"`
	List  []OrderInfo `json:"list"`
}

type ChangeFood struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Describe string `json:"describe"`
	Cost     int    `json:"cost"`
	Foodless int    `json:"foodless"`
}

func (u RestService) Rest(msg string) string {
	return fmt.Sprintf("your name is %v", msg)
}

func (u RestService) Restinfo(c *gin.Context) (resp *Restinforesp, err error) {
	session := sessions.Default(c)
	fmt.Println(session.Get("userId"), reflect.TypeOf(session.Get("userId")))
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return nil, err
	}
	return &Restinforesp{
		UserId: rest.ID,
		Number: rest.Name,
		Status: rest.Status,
	}, nil
}

func (u RestService) HandleFoodAddservice(c *gin.Context) (resp *Foodinforesp, err error) {
	var form struct {
		Foodname string `form:"foodname" json:"foodname" binding:"required"`
		Foodless int    `json:"foodless" form:"foodless" binding:"required"`
		Cost     int    `json:"cost" form:"cost" binding:"required"`
		Describe string `json:"describe" form:"describe" binding:"required"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return nil, err
	}
	session := sessions.Default(c)
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return nil, err
	}
	food := &model.Food{}
	if err = model.DB.Where("foodname = ?", form.Foodname).First(&food).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		food = &model.Food{
			Foodless: form.Foodless,
			Foodname: form.Foodname,
			Cost:     form.Cost,
			Describe: form.Describe,
			ResID:    rest.ID,
		}
		if err = model.DB.Create(&food).Error; err != nil {
			return nil, err
		}
		return &Foodinforesp{
			ID:       food.ID,
			Foodname: food.Foodname,
			Cost:     food.Cost,
			Describe: food.Describe,
		}, nil
	}
	return nil, errors.New("已存在相同的菜品")
}

func (u RestService) HandleNewRestservice(c *gin.Context) (resp int, err error) {
	var form struct {
		Name        string `form:"name" json:"name" binding:"required"`
		Address     string `json:"address" form:"address" binding:"required"`
		LicenceStar int    `json:"licencestar" form:"licencestar" binding:"required"`
		Describe    string `json:"describe" form:"describe" binding:"required"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return -1, err
	}
	session := sessions.Default(c)
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err == nil {
		return -1, errors.New("店铺正在审核中或已经开过店,请访问相关api查询店铺状况")
	}
	rest = &model.Restaurant{
		Name:        form.Name,
		Address:     form.Address,
		UserID:      session.Get("userId").(int),
		LicenceStar: form.LicenceStar,
		Describe:    form.Describe,
		Status:      0,
	}
	if err = model.DB.Create(&rest).Error; err != nil {
		return -1, err
	}
	return rest.Status, nil
}

func (u RestService) HandleNewRestinfoService(c *gin.Context) (re string, err error) {
	session := sessions.Default(c)
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return "", err
	}
	switch rest.Status {
	case 0:
		re = "正在审核中"
	case 1:
		re = "当前正在营业(审核已通过)"
	case 2:
		re = "审核未通过"
	case 3:
		re = "当前正在休息"
	}
	return re, nil
}

func (u RestService) HandleCheckService(c *gin.Context) (resp *Orderresp, err error) {
	var uri struct {
		Page  int `form:"page" json:"page" uri:"page" binding:"required" validate:"gte:1"`
		Limit int `form:"limit" json:"limit" uri:"limit" binding:"required"`
	}

	if err := c.ShouldBindQuery(&uri); err != nil {
		return nil, err
	}
	session := sessions.Default(c)
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return nil, err
	}
	offest := (uri.Page - 1) * uri.Limit
	restList := &[]OrderInfo{}
	var total int64
	tx := model.DB.Model(&model.Order{})
	tx.Limit(uri.Limit).Offset(offest)
	if err := tx.Where("resid = ? AND status != ?", rest.ID, 3).Find(restList).Count(&total).Error; err != nil {
		return nil, err
	}

	return &Orderresp{
		Total: total,
		List:  *restList,
	}, nil
}

func (u RestService) HandleChangeService(c *gin.Context) (err error) {
	var form struct {
		ID     int `json:"id" form:"id" binding:"required"`
		Status int `json:"status" form:"status" binding:"required"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return err
	}
	session := sessions.Default(c)
	order := &model.Order{}
	if err = model.DB.Where("id = ?", form.ID).First(&order).Error; err != nil {
		return err
	}
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return err
	}
	if order.Resid != rest.ID {
		return errors.New("请不要修改非本店的订单")
	}
	if order.Status == 3 {
		return errors.New("此份订单已结单")
	}
	switch form.Status {
	case 1:
		return errors.New("不能改为未接单")
	case 2:
		order.Status = 2
	case 3:

		order.Status = 3
		food := &model.Food{}
		model.DB.Where("id = ?", order.Foodid).First(&food)
		food.Foodless = food.Foodless - 1
		model.DB.Save(&food)
	default:
		return errors.New("参数输入错误")
	}
	model.DB.Save(&order)
	return nil
}

func (u RestService) HandleChangeFoodService(c *gin.Context) (resp *ChangeFood, err error) {
	var form struct {
		ID       int    `json:"id" form:"id" binding:"required"`
		Cost     int    `json:"cost" form:"cost"`
		Describe string `json:"describe" form:"describe"`
		Foodless int    `json:"foodless" form:"foodless"`
		Name     string `json:"name" form:"name"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return nil, err
	}
	session := sessions.Default(c)
	food := &model.Food{}
	if err = model.DB.Where("id = ?", form.ID).First(&food).Error; err != nil {
		return nil, err
	}
	rest := &model.Restaurant{}
	if err = model.DB.Where("userid = ?", session.Get("userId").(int)).First(&rest).Error; err != nil {
		return nil, err
	}
	if food.ResID != rest.ID {
		return nil, errors.New("请不要修改非本店的商品")
	}
	flag := 0
	if form.Cost != 0 && food.Cost != form.Cost {
		food.Cost = form.Cost
		flag = 1
	}
	if form.Describe != "" && food.Describe != form.Describe {
		food.Describe = form.Describe
		fmt.Println(form.Describe)
	}
	if form.Foodless != 0 && food.Foodless != form.Foodless {
		food.Foodless = form.Foodless
		flag = 1
	}
	if form.Name != "" && food.Foodname != form.Name {
		food.Foodname = form.Name
		flag = 1
	}
	if flag == 1 {
		food.Like = 0
		food.Number = 0
		order := &model.Order{}
		if err := model.DB.Where("foodid = ?", form.ID).Delete(&order).Error; err != nil {
			return nil, err
		}
	}
	model.DB.Save(&food)
	return &ChangeFood{
		ID:       food.ID,
		Name:     food.Foodname,
		Describe: food.Describe,
		Cost:     food.Cost,
		Foodless: food.Foodless,
	}, nil
}
