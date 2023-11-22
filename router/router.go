package router

import (
	"net/http"
	"path/filepath"

	"cinema.com/controller"
	"cinema.com/middleware"
	"cinema.com/repository"
	"github.com/gin-gonic/gin"
)

func NewRouter(userRepository repository.UserRepository, authenticationController *controller.AuthenticationController, filmController *controller.FilmController) *gin.Engine {
	service := gin.Default()

	service.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, "welcome home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	service.GET("/image/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		imagePath := filepath.Join("./uploads", filename)

		c.File(imagePath)
	})

	router := service.Group("/api")
	{
		authenticationRouter := router.Group("/authentication")
		authenticationRouter.POST("/register", authenticationController.Register)
		authenticationRouter.POST("/login", authenticationController.Login)
	}

	usersRouter := router.Group("/films")
	{
		middelware := middleware.DeserilizeUser(userRepository)
		
		usersRouter.GET("", middelware, filmController.GetFilm)
		usersRouter.POST("", middelware, filmController.CreateFilm)
		usersRouter.DELETE("/remove", middelware, filmController.DeleteFilm)
		usersRouter.GET("/seats", middelware, filmController.GetSeat)
		usersRouter.POST("/booking", middelware, filmController.BookingFilm)
		usersRouter.GET("actor", middelware, filmController.GetActor)
		usersRouter.POST("actor", middelware, filmController.AddActor)
		usersRouter.DELETE("/actor/delete", middelware, filmController.DeleteActor)
		usersRouter.POST("/connect", middelware, filmController.ConnectActor)
	}

	return service
}
