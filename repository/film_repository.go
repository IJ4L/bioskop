package repository

import (
	"cinema.com/data/response"
	"cinema.com/model"
)

type FilmRepository interface {
	GetAll() []model.Film
	GetSeat(id int) []response.SeatStatus
}
