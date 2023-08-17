package router

import (
	"net/http"

	"cinema.com/controller"
	"cinema.com/middleware"
	"cinema.com/repository"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UserRepository, authenticationController *controller.AuthenticationController, filmController *controller.FilmController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	{
		authenticationRouter := router.Group("/authentication")
		authenticationRouter.POST("/register", authenticationController.Register)
		authenticationRouter.POST("/login", authenticationController.Login)
	}

	usersRouter := router.Group("/films")
	{
		usersRouter.GET("", middleware.DeserilizeUser(userRepository), filmController.GetFilm)
		usersRouter.GET("/seats", middleware.DeserilizeUser(userRepository), filmController.GetSeat)
	}

	return service
}
