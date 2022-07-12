package pkg

import (
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/base"
	"time"
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