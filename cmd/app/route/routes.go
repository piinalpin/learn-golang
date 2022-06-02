package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	var router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	return router
}