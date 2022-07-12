package base

import (
	"time"
)

type BaseDto struct {
	CreatedBy	string		`json:"created_by"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	*time.Time	`json:"updated_at"`
	DeletedAt	*time.Time	`json:"deleted_at"`
}
