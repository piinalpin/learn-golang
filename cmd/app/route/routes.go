package route

import (
	"learn-rest-api/cmd/app/middleware"
	"learn-rest-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(init *config.Initialization) *gin.Engine {
	var router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	var auth = router.Group("/auth")
	{
		auth.POST("/login", init.AuthCtrl.Login)
		auth.POST("/refresh", init.AuthCtrl.RefreshToken)
	}

	var api = router.Group("/api")
	{

		var author = api.Group("/author").Use(middleware.Authorization(init.TokenUtil))
		{
			author.GET("", init.AuthorCtrl.GetAllAuthor)
			author.POST("", init.AuthorCtrl.CreateAuthor)
		}

	}

	return router
}
