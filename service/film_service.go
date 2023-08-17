package service

import (
	"cinema.com/data/response"
	"cinema.com/model"
)

type FilmService interface {
	GetAllFilm() ([]model.Film, error)
	GetSeat(id int) ([]response.SeatStatus, error)
}