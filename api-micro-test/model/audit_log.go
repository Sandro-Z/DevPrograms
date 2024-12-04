package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditLog struct {
	ID          string `gorm:"type:CHAR(36) NOT NULL;index;comment:ID"`
	From        string `gorm:"type:CHAR(36) NOT NULL;comment:发件人"`
	To          string `gorm:"type:VARCHAR(255) NOT NULL;comment:收件人"`
	Subject     string `gorm:"type:VARCHAR(255) NOT NULL;comment:主题"`
	ContentBody string `gorm:"type:TEXT NOT NULL;comment:信息"`
	ContentType string `gorm:"type:VARCHAR(16) NOT NULL;comment:信息"`

	BaseModel
}

func (AuditLog) TableName() string {
	t := time.Now().Local()
	return fmt.Sprintf("audit_logs_%04d%02d", t.Year(), t.Month())
}

func (e *AuditLog) BeforeCreate(_ *gorm.DB) (err error) {
	if e.ID == "" {
		e.ID = uuid.NewString()
	}
	return nil
}

func (e *AuditLog) AfterCreate(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) BeforeUpdate(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) AfterUpdate(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) BeforeSave(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) AfterSave(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) BeforeDelete(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) AfterDelete(_ *gorm.DB) (err error) {
	return nil
}

func (e *AuditLog) AfterFind(_ *gorm.DB) error {
	return nil
}
