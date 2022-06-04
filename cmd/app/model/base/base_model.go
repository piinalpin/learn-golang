package base

import (
	"time"
)

type BaseModel struct {
	CreatedBy	string		`json:"created_by" gorm:"column:created_by; size:255; not null; default:'SYSTEM'"`
	CreatedAt	time.Time	`json:"created_at" gorm:"column:created_at; autoCreateTime; not null"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"column:updated_at; autoUpdateTime; null"`
	DeletedAt	time.Time	`json:"deleted_at" gorm:"column:deleted_at; null"`
}