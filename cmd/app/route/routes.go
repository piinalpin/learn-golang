package route

import (
	"learn-rest-api/cmd/app/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
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
			var authorController = new (controller.AuthorController)
			author.GET("", authorController.GetAllAuthor)
		}

	}

	return router
}