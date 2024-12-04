package model

type Order struct {
	BaseModel

	Userid   int    `gorm:"column:userid;not null;comment:'点菜用户ID'" json:"userid"`
	Resid    int    `gorm:"column:resid;not null;comment:'商家ID'" json:"resid"`
	Foodid   int    `gorm:"column:foodid;not null;comment:'菜品id'" json:"foodid"`
	Foodname string `gorm:"column:foodname;not null;comment:'菜品名'" json:"foodname"`
	Status   int    `gorm:"column:status;not null;comment:'当前订单状况.1为未接单,2为已接单,3为已结单'" json:"status"`
}

func (Order) TableName() string {
	return "orders"
}
