package controller

import (
	"net/http"

	"cinema.com/data/request"
	"cinema.com/data/response"
	"cinema.com/service"
	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationService service.AuthenticationService
	ctx                   *gin.Context
}

func NewAuthenticationController(authenticationService service.AuthenticationService, ctx *gin.Context) *AuthenticationController {
	return &AuthenticationController{authenticationService: authenticationService, ctx: ctx}
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid email or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	token, err_token := c.authenticationService.Login(loginRequest)

	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid email or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AuthenticationController) Register(ctx *gin.Context) {
	createUserRequest := request.CreateUsersRequest{}
	ctx.ShouldBindJSON(&createUserRequest)

	_, err := c.authenticationService.Register(createUserRequest)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: err.Error(),
		})
		ctx.Abort()
		return
	}

	webRespones := response.Response{
		Code:    200,
		Status:  "OK",
		Message: "Berhasil membuat akun",
	}

	ctx.JSON(http.StatusOK, webRespones)
}
