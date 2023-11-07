package main

import (
	"net/http"
	"time"

	"cinema.com/config"
	"cinema.com/controller"
	"cinema.com/helper"
	"cinema.com/repository"
	"cinema.com/router"
	"cinema.com/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	db := config.InitDatabase()
	validate := validator.New()
	var ctx *gin.Context

	
	// db.AutoMigrate(&model.Film{}, &model.Actor{}, &model.Seat{}, &model.Booking{}, &model.User{}, &model.Bloc{}, &model.FilmActor{})

	userRepository := repository.NewUsersRepositoryImpl(db)
	authenticationService := service.NewAuthenticationService(userRepository, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService, ctx)

	filmRepository := repository.NewFilmRepositoryImpl(db)
	filmService := service.NewFilmService(filmRepository, validate)
	filmController := controller.NewFilmController(filmService, ctx)

	routes := router.NewRouter(userRepository, authenticationController, filmController)

	server := &http.Server{
		Addr:           "localhost:8080",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
