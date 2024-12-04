package service

import (
	"errors"
	"exam_go/model"
	"fmt"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
}

type Userinforesp struct {
	UserId int    `json:"userId"`
	Number string `json:"number"`
	Name   string `json:"name"`
}

type SearchRestsResponse struct {
	Total int64       `json:"total"`
	List  []RestsList `json:"list"`
}

type RestsList struct {
	ID      int    `json:"ID,omitempty"`
	Name    string `json:"name,omitempty"`
	Status  string `json:"status,omitempty"`
	People  string `json:"people,omitempty"`
	Address string `json:"address,omitempty"`
}

type SearchFoodsResponse struct {
	Total int64       `json:"total"`
	List  []FoodsList `json:"list"`
}

type FoodsList struct {
	ID       int    `json:"ID,omitempty"`
	Foodname string `json:"foodname,omitempty"`
	Cost     int    `json:"cost,omitempty"`
	Number   string `json:"number,omitempty"`
	Like     string `json:"like,omitempty"`
	Foodless string `json:"foodless,omitempty"`
}

type Orderfoodres struct {
	Status int `json:"status,omitempty"`
	ID     int `json:"titleid,omitempty"`
}

func (u UserService) User(msg string) string {
	return fmt.Sprintf("your name is %v", msg)
}

func (u UserService) Userinfo(c *gin.Context) (resp *Userinforesp, err error) {
	session := sessions.Default(c)
	fmt.Println(session.Get("userId"), reflect.TypeOf(session.Get("userId")))
	user := &model.User{}
	if err = model.DB.Where("id = ?", session.Get("userId").(int)).First(&user).Error; err != nil {
		return nil, err
	}
	return &Userinforesp{
		UserId: user.ID,
		Number: user.Number,
		Name:   user.Name,
	}, nil
}

func (u UserService) HandleSignupservice(c *gin.Context) (resp *Userinforesp, err error) {
	var form struct {
		Number   string `form:"number" json:"number" binding:"required"`
		Name     string `json:"name" form:"name" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return nil, err
	}
	user := &model.User{}
	if err = model.DB.Where("number = ?", form.Number).First(&user).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		user = &model.User{
			Number:   form.Number,
			Name:     form.Name,
			Password: form.Password,
			Status:   0,
		}
		if err = model.DB.Create(&user).Error; err != nil {
			return nil, err
		}
		session := sessions.Default(c)
		session.Set("userId", user.ID)
		session.Set("status", 0)
		if err = session.Save(); err != nil {
			return nil, err
		}
		return &Userinforesp{
			UserId: user.ID,
			Number: user.Number,
			Name:   user.Name,
		}, nil
	}
	return nil, errors.New("已存在相同的学工号")
}

func (u UserService) HandleLoginService(c *gin.Context) (resp *Userinforesp, err error) {
	var form struct {
		Number   string `form:"number" json:"number" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
	if err = c.ShouldBind(&form); err != nil {
		return nil, err
	}
	user := &model.User{}
	if err = model.DB.Where("number = ?", form.Number).First(&user).Error; err != nil {
		return nil, err
	}
	if form.Password != user.Password {
		return nil, errors.New("密码不正确")
	}
	session := sessions.Default(c)
	session.Set("userId", user.ID)
	session.Set("status", user.Status)
	// session.Options(sessions.Options{MaxAge: 30})
	if err = session.Save(); err != nil {
		return nil, err
	}
	fmt.Println(session.Get("userId"))
	return &Userinforesp{
		UserId: user.ID,
		Name:   user.Name,
		Number: user.Number,
	}, nil
}

func (u UserService) HandleGetRestListService(c *gin.Context) (resp *SearchRestsResponse, err error) {
	var uri struct {
		Page  int `form:"page" json:"page" uri:"page" binding:"required" validate:"gte:1"`
		Limit int `form:"limit" json:"limit" uri:"limit" binding:"required"`
	}

	if err := c.ShouldBindQuery(&uri); err != nil {
		return nil, err
	}
	resp = &SearchRestsResponse{}
	println(uri.Page)
	println(uri.Limit)
	offest := (uri.Page - 1) * uri.Limit
	restList := &[]RestsList{}
	var total int64
	tx := model.DB.Model(&model.Restaurant{}).Count(&total)
	tx.Limit(uri.Limit).Offset(offest)
	if err := tx.Find(restList).Error; err != nil {
		return nil, err
	}
	resp.Total = total
	resp.List = *restList
	return resp, nil
}

func (u UserService) HandleGetRestFoodListService(c *gin.Context) (resp *SearchFoodsResponse, err error) {
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

func (u UserService) HandleGetFoodInfoService(c *gin.Context) (resp *[]FoodsList, err error) {
	var uri struct {
		ID int `uri:"id"`
	}

	if err := c.ShouldBindUri(&uri); err != nil {
		return nil, err
	}
	foodList := &[]FoodsList{}
	tx := model.DB.Model(&model.Food{})
	if err := tx.Where("id = ?", uri.ID).Find(foodList).Error; err != nil {
		return nil, err
	}
	return foodList, nil
}

func (u UserService) HandleOrderFoodService(c *gin.Context) (resp *Orderfoodres, err error) {
	var uri struct {
		ID int `form:"id" json:"id" binding:"required"`
	}
	session := sessions.Default(c)
	if err := c.ShouldBind(&uri); err != nil {
		return nil, err
	}
	orderfood := &Orderfoodres{}
	tx := &model.Food{}
	if err := model.DB.Where("id = ?", uri.ID).First(&tx).Error; err != nil {
		return nil, err
	}
	rest := &model.Restaurant{}
	if err = model.DB.Where("id = ?", tx.ResID).First(&rest).Error; err != nil {
		return nil, err
	}
	if rest.Status != 1 || tx.Foodless == 0 {
		return nil, errors.New("商家未开店或商品已售罄")
	}
	order := &model.Order{
		Userid:   session.Get("userId").(int),
		Resid:    rest.ID,
		Foodid:   tx.ID,
		Foodname: tx.Foodname,
		Status:   1,
	}
	if err = model.DB.Create(&order).Error; err != nil {
		return nil, err
	}
	orderfood.ID = order.ID
	orderfood.Status = 1
	return orderfood, nil
}

func (u UserService) HandleGetOrderInfoService(c *gin.Context) (resp int, err error) {
	var uri struct {
		ID int `uri:"id"`
	}

	session := sessions.Default(c)
	if err := c.ShouldBindUri(&uri); err != nil {
		return -1, err
	}
	tx := &model.Order{}
	if err := model.DB.Where("id = ?", uri.ID).First(&tx).Error; err != nil {
		return -1, err
	}
	if tx.Userid != session.Get("userId").(int) {
		return -1, errors.New("查询的不是你的订单哦")
	}
	return tx.Status, nil
}

func (u UserService) HandleLikeFood(c *gin.Context) (err error) {
	var uri struct {
		Like int `form:"like" json:"like" binding:"required"`
		ID   int `form:"id" json:"id" binding:"required"`
	}
	session := sessions.Default(c)
	if err := c.ShouldBind(&uri); err != nil {
		return err
	}
	if uri.Like > 0 {
		uri.Like = 1
		fmt.Println("turrr")
	}
	order := &model.Order{}
	if err := model.DB.Where("id = ?", uri.ID).First(&order).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if order.Status != 3 {
		return errors.New("当前未结单,请吃完后再评价")
	}
	if order.Userid != session.Get("userId").(int) {
		return errors.New("请不要评价其他人的订单")
	}
	tx := &model.Food{}
	if err := model.DB.Where("id = ?", order.Foodid).First(&tx).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	restt := model.DB.Model(&model.Food{})
	if err := restt.Where("id = ?", uri.ID).Update("like", tx.Like+uri.Like).Error; err != nil {
		return err
	}
	return nil
}
