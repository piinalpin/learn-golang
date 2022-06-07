package route

import (
	"learn-rest-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(init *config.Initialization) *gin.Engine {
	var router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	var api = router.Group("/api")
	{
		
		api.GET("", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		})

		var author = api.Group("/author")
		{
			author.GET("", init.AuthorCtrl.GetAllAuthor)
			author.POST("", init.AuthorCtrl.CreateAuthor)
		}

	}

	return router
}