package pkg

import (
	"time"
	_ "github.com/gin-gonic/gin"
	"learn-rest-api/cmd/app/model/base"
	"learn-rest-api/cmd/app/constant"
)

// func BuildResponse[T any](responseKey constant.ResponseKey, data T, httpStatus int) (c *gin.Context) {
// 	c.JSON(httpStatus, build(responseKey.GetKey(), responseKey.GetMessage(), data))
// 	return
// }

func BuildResponse[T any](responseKey constant.ResponseKey, data T) base.ApiResponse[T] {
	return build(responseKey.GetKey(), responseKey.GetMessage(), data)
}

func build[T any](responseKey string, message string, data T) base.ApiResponse[T] {
	return base.ApiResponse[T]{
		Timestamp: time.Now(),
		ResponseKey:    responseKey,
		Message:   message,
		Data:      data,
	}
}