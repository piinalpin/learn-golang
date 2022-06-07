package pkg

import (
	"time"
	"learn-rest-api/cmd/app/model/base"
	"learn-rest-api/cmd/app/constant"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseKey respkey.ResponseKey, data T) base.ApiResponse[T] {
	return BuildResponse_(responseKey.GetKey(), responseKey.GetMessage(), data)
}

func BuildResponse_[T any](responseKey string, message string, data T) base.ApiResponse[T] {
	return build(responseKey, message, data)
}

func build[T any](responseKey string, message string, data T) base.ApiResponse[T] {
	return base.ApiResponse[T]{
		Timestamp: time.Now(),
		ResponseKey:    responseKey,
		Message:   message,
		Data:      data,
	}
}