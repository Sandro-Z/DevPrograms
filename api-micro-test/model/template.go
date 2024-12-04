package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Template struct {
	ID string `gorm:"type:CHAR(36) NOT NULL;uniqueIndex;comment:模板ID"`

	Subject     string `gorm:"type:VARCHAR(255) NOT NULL;comment:主题"`
	Importance  string `gorm:"type:VARCHAR(32) NOT NULL;index;comment:重要性"`
	ContentBody string `gorm:"type:VARCHAR(5000) NOT NULL;comment:内容模版"`
	ContentType string `gorm:"type:VARCHAR(16) NOT NULL;comment:内容类型"`

	UserID string `gorm:"type:CHAR(36) NOT NULL;index;comment:用户ID"`
	BaseModel
}

func (Template) TableName() string {
	return "templates"
}

func (e *Template) BeforeCreate(_ *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}
	return nil
}
