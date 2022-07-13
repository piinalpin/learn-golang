package route

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/middleware"
	"learn-rest-api/config"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Router(init *config.Initialization) *gin.Engine {
	godotenv.Load()

	allowedOrigins := strings.Split(os.Getenv("application.cors.allowed-origins"), ",")
	allowedMethods := strings.Split(os.Getenv("application.cors.allowed-methods"), ",")
	allowedHeaders := strings.Split(os.Getenv("application.cors.allowed-headers"), ",")
	exposedHeaders := strings.Split(os.Getenv("application.cors.exposed-headers"), ",")

	var router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins:  allowedOrigins,
		AllowMethods:  allowedMethods,
		AllowHeaders:  allowedHeaders,
		ExposeHeaders: exposedHeaders,
		AllowCredentials: true,
		MaxAge: 1 * time.Hour,
	}))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	var auth = router.Group("/auth").Use(middleware.Basic())
	{
		auth.POST("/login", init.AuthCtrl.Login)
		auth.POST("/refresh", init.AuthCtrl.RefreshToken)
	}

	var api = router.Group("/api")
	{

		var author = api.Group("/author").Use(middleware.Bearer(init.TokenUtil))
		author.Use(middleware.RoleBasedAccessControl(constant.Admin.GetRole()))
		{
			author.GET("", init.AuthorCtrl.GetAllAuthor)
			author.POST("", init.AuthorCtrl.CreateAuthor)
		}

		var user = api.Group("/user").Use(middleware.Bearer(init.TokenUtil))
		{
			user.GET("/me", init.UserCtrl.Me)
		}

	}

	return router
}
