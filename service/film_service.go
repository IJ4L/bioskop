package service

import (
	"cinema.com/data/request"
	"cinema.com/data/response"
	"cinema.com/model"
)

type FilmService interface {
	GetAllFilm() ([]model.Film, error)
	GetSeat(id int) ([]response.SeatStatus, error)
	AddFilm(film request.AddFilm) error
	DeleteFilm(id uint) error
	BookingFilm(booking request.BookingFilm) (request.BookingFilm, error)
	GetActor() []model.Actor
	AddActor(actor request.AddActor) error
	DeleteActor(id uint) error
	ConnectActor(request.ConnectActor) error
}
