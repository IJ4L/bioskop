package repository

import (
	"cinema.com/data/response"
	"cinema.com/model"
	"gorm.io/gorm"
)

type FilmRepositoryImpl struct {
	DB *gorm.DB
}

func NewFilmRepositoryImpl(db *gorm.DB) *FilmRepositoryImpl {
	return &FilmRepositoryImpl{DB: db}
}

func (r *FilmRepositoryImpl) GetAll() []model.Film {
	films := []model.Film{}
	r.DB.Preload("Actors").Find(&films)

	return films
}

func (r *FilmRepositoryImpl) GetSeat(id int) []response.SeatStatus {
	results := []response.SeatStatus{}

	r.DB.Table("blocs").
		Select("seats.seat_number, CASE WHEN bookings.id_pemesan IS NOT NULL THEN 'Sold' ELSE 'Ready' END AS StatusKursi").
		Joins("JOIN seats ON blocs.id_bloc = seats.id_bloc").
		Joins("LEFT JOIN bookings ON seats.id_seat = bookings.id_seat AND bookings.id_film = ?", id).
		Scan(&results)

	return results
}
