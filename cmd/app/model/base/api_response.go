package base

import (
	"time"
)

type ApiResponse[T any] struct {
	Timestamp		time.Time		`json:"timestamp"`
	ResponseKey		string			`json:"response_key"`
	Message			string			`json:"message"`
	Data			T				`json:"data"`
}