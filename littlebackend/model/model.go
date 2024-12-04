package model

import "time"

type BaseModel struct {
	ID        int       `gorm:"type:BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey;comment:主键" json:"id"`
	CreatedAt time.Time `gorm:"comment:'创建时间';not null;column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"comment:'更新时间';not null;column:updated_at" json:"updated_at"`
}
