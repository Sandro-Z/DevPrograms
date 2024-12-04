package model

type User struct {
	BaseModel

	Name     string `gorm:"column:name;not null;comment:'姓名'" json:"name"`
	Password string `gorm:"column:password;not null;comment:'密码'" json:"password"`
	Number   string `gorm:"column:number;not null;comment:'学工号'" json:"number"`
	Status   int    `gorm:"column:status;not null;comment:'状态 0-普通用户 1-管理员';type:tinyint(2)" json:"status"`
}

func (User) TableName() string {
	return "users"
}
