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

var (
	ctx *gin.Context
)

func main() {
	db := config.InitDatabase()
	validate := validator.New()

	userRepository := repository.NewUsersRepositoryImpl(db)
	authenticationService := service.NewAuthenticationService(userRepository, validate)
	authenticationController := controller.NewAuthenticationController(authenticationService, ctx)

	filmRepository := repository.NewFilmRepositoryImpl(db)
	filmService := service.NewFilmService(filmRepository, validate)
	filmController := controller.NewFilmController(filmService, ctx)

	routes := router.NewRouter(userRepository, authenticationController, filmController)

	server := &http.Server{
		Addr:           "localhost:3000",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
