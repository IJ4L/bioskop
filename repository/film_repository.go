package repository

import (
	"cinema.com/data/request"
	"cinema.com/data/response"
	"cinema.com/model"
)

type FilmRepository interface {
	GetAll() []model.Film
	GetSeat(id int) []response.SeatStatus
	AddFilm(film request.AddFilm) error
	DeleteFilm(id uint) error
	BookingFilm(booking request.BookingFilm) (request.BookingFilm, error)
	FindById(id uint) error
}
