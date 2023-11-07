package controller

import (
	"net/http"
	"strconv"

	"cinema.com/data/request"
	"cinema.com/data/response"
	"cinema.com/service"
	"cinema.com/utils"
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

func (controller *FilmController) CreateFilm(ctx *gin.Context) {
	// Upload the image (poster)
	poster, err := utils.UploadImage(ctx, "poster")
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Retrieve the form data
	form, err := ctx.MultipartForm()
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Perform validation to ensure required fields are not empty
	requiredFields := []string{"judul", "deskripsi", "genre", "imdb", "durations", "views", "p-g", "price", "show_date", "show_times"}
	for _, field := range requiredFields {
		if len(form.Value[field]) == 0 {
			webResponse := response.ErorrResponse{
				Code:   http.StatusBadRequest,
				Status: "Bad Request",
			}
			ctx.JSON(http.StatusBadRequest, webResponse)
			return
		}
	}

	// Create the AddFilm struct using the form data
	film := request.AddFilm{
		Judul:     form.Value["judul"][0],
		Desk:      form.Value["deskripsi"][0],
		Genre:     form.Value["genre"][0],
		Imdb:      form.Value["imdb"][0],
		Poster:    poster,
		Durations: form.Value["durations"][0],
		Views:     form.Value["views"][0],
		Pg:        form.Value["p-g"][0],
		Price:     form.Value["price"][0],
		ShowDate:  form.Value["show_date"][0],
		ShowTimes: form.Value["show_times"][0],
	}

	// Add the film using the filmService
	if err := controller.filmService.AddFilm(film); err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Return a success response
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success add film",
		Data:    film,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) DeleteFilm(ctx *gin.Context) {
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

	if err := controller.filmService.DeleteFilm(uint(id)); err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success delete film",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) BookingFilm(ctx *gin.Context) {
	var booking request.BookingFilm
	err := ctx.ShouldBindJSON(&booking)
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	_, err = controller.filmService.BookingFilm(booking)
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success booking film",
		Data:    booking,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *FilmController) GetActor(ctx *gin.Context) {
	actor := controller.filmService.GetActor()
	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success get actor",
		Data:    actor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) AddActor(ctx *gin.Context) {
	photo, err := utils.UploadImage(ctx, "photo")
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "1",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Retrieve the form data
	form, err := ctx.MultipartForm()
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "2",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Perform validation to ensure required fields are not empty
	requiredFields := []string{"name", "profesion"}
	for _, field := range requiredFields {
		if len(form.Value[field]) == 0 {
			webResponse := response.ErorrResponse{
				Code:   http.StatusBadRequest,
				Status: "3",
			}
			ctx.JSON(http.StatusBadRequest, webResponse)
			return
		}
	}

	// Create the AddFilm struct using the form data
	actor := request.AddActor{
		Name:      form.Value["name"][0],
		Photo:     photo,
		Profesion: form.Value["profesion"][0],
	}

	// Add the film using the filmService
	if err := controller.filmService.AddActor(actor); err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "4",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success add actor",
		Data:    actor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) DeleteActor(ctx *gin.Context) {
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

	if err := controller.filmService.DeleteActor(uint(id)); err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success delete actor",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FilmController) ConnectActor(ctx *gin.Context) {
	var connect request.ConnectActor
	err := ctx.ShouldBindJSON(&connect)
	if err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad 1",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	if err := controller.filmService.ConnectActor(connect); err != nil {
		webResponse := response.ErorrResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		ctx.Abort()
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Success connect actor",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
