package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"cinema.com/data/response"
	"cinema.com/service"
	"github.com/gin-gonic/gin"
)

type FilmController struct {
	filmService service.FilmService
	ctx         *gin.Context
}

func NewFilmController(filmService service.FilmService, ctx *gin.Context) *FilmController {
	return &FilmController{filmService: filmService, ctx: ctx}
}

func (controller *FilmController) GetFilm(ctx *gin.Context) {
	film, err := controller.filmService.GetAllFilm()

	if err != nil {
		webResponse := response.Response{
			Code:   400,
			Status: "Failed",
			Data:   nil,
		}

		ctx.JSON(http.StatusInternalServerError, webResponse)
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Success get all film",
		Data:    film,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) GetSeat(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid id",
			Data:    nil,
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	fmt.Println("ID:", id)

	seat, err := controller.filmService.GetSeat(id)
	if err != nil {
		webResponse := response.Response{
			Code:    400,
			Status:  "Failed",
			Message: "Failed get seat",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Success get seat",
		Data:    seat,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
