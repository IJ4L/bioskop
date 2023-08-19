package repository

import (
	"cinema.com/data/request"
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

func (repo *FilmRepositoryImpl) GetAll() []model.Film {
	films := []model.Film{}
	repo.DB.Preload("Actors").Find(&films)

	return films
}

func (repo *FilmRepositoryImpl) GetSeat(id int) []response.SeatStatus {
	results := []response.SeatStatus{}

	repo.DB.Table("blocs").
		Select("seats.seat_number, CASE WHEN bookings.id_pemesan IS NOT NULL THEN 'Sold' ELSE 'Ready' END AS StatusKursi").
		Joins("JOIN seats ON blocs.id_bloc = seats.id_bloc").
		Joins("LEFT JOIN bookings ON seats.id_seat = bookings.id_seat AND bookings.id_film = ?", id).
		Scan(&results)

	return results
}

func (repo *FilmRepositoryImpl) AddFilm(film request.AddFilm) error {

	db := repo.DB.Table("films").Create(&film)
	if db.Error != nil {
		err := db.Error
		return err
	}
	return nil
}

func (repo *FilmRepositoryImpl) DeleteFilm(id uint) error {
	db := repo.DB.Delete(&model.Film{}, id)
	if db.Error != nil {
		err := db.Error
		return err
	}
	return nil
}

func (repo *FilmRepositoryImpl) BookingFilm(booking request.BookingFilm) (request.BookingFilm, error) {
	db := repo.DB.Table("bookings")
	result := db.Save(&booking)
	if result.Error != nil {
		return booking, result.Error
	}
	return booking, nil
}

func (repo *FilmRepositoryImpl) FindById(id uint) error {
	db := repo.DB.First(&model.Film{}, id)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
