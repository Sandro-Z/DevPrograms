package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mail struct {
	ID string `gorm:"type:CHAR(36) NOT NULL;uniqueIndex;comment:邮件ID"`
	To string `gorm:"type:VARCHAR(255) NOT NULL;index;comment:收件人"`

	Fields Fields `gorm:"type:TEXT NOT NULL;comment:内容"`
	Status string `gorm:"type:VARCHAR(32) NOT NULL;index;comment:状态"`

	Template *Template `gorm:"-"`

	TmplID string `gorm:"type:CHAR(36) NOT NULL;index;comment:模版ID"`
	UserID string `gorm:"type:CHAR(36) NOT NULL;index;comment:用户ID"`
	BaseModel
}

func (Mail) TableName() string {
	return "mails"
}

func (e *Mail) BeforeCreate(_ *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}
	return nil
}
