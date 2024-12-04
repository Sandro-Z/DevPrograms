package model

type Food struct {
	BaseModel

	ResID    int    `gorm:"column:resid;not null;comment:'商家ID'" json:"resid"`
	Foodname string `gorm:"column:foodname;not null;comment:'菜品名'" json:"foodname"`
	Cost     int    `gorm:"column:cost;not null;comment:'价格';" json:"cost"`
	Number   int    `gorm:"column:number;not null;comment:'点过的人数';" json:"number"`
	Like     int    `gorm:"column:like;not null;comment:'好评数';" json:"like"`
	DisLike  int    `gorm:"column:dislike;not null;comment:'差评数';" json:"dislike"`
	Foodless int    `gorm:"column:foodless;not null;comment:'剩余数量';" json:"foodless"`
	Describe string `gorm:"column:describe;not null;comment:'菜品描述'" json:"describe"`
}

func (Food) TableName() string {
	return "foods"
}

type Restaurant struct {
	BaseModel

	Name        string `gorm:"column:name;not null;comment:'商家名'" json:"name"`
	UserID      int    `gorm:"column:userid;not null;comment:'开店用户的id'" json:"userid"`
	Status      int    `gorm:"column:status;not null;comment:'开店状况.0审核中,1开店,2审核未通过,3闭店';type:tinyint(4)" json:"status"`
	Address     string `gorm:"column:address;not null;comment:'地址'" json:"address"`
	LicenceStar int    `gorm:"column:licenceStar;not null;comment:'星级(1-5)';type:tinyint(6)" json:"licenceStar"`
	People      int    `gorm:"column:people;not null;comment:'用餐总人数';" json:"people"`
	Describe    string `gorm:"column:describe;not null;comment:'商家介绍'" json:"describe"`

	// Foods []Food `gorm:"foreignKey:ID;references:ResID"`
}

func (Restaurant) TableName() string {
	return "rests"
}
