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
	result := db.Create(&booking)
	if result.Error != nil {
		return booking, result.Error
	}
	return booking, nil
}

func (repo *FilmRepositoryImpl) FindById(id uint) error {
	var film model.Film
	err := repo.DB.First(&film, id).Error
	if err != nil {
		return err
	}
	return nil
}

// AddActor implements FilmRepository.
func (repo *FilmRepositoryImpl) AddActor(actor request.AddActor) error {
	db := repo.DB.Table("actors").Create(&actor)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

// ConnectActor implements FilmRepository.
func (repo *FilmRepositoryImpl) ConnectActor(connect request.ConnectActor) error {
	db := repo.DB.Table("film_actors").Create(&connect)

	if db.Error != nil {
		return db.Error
	}

	return nil
}

// DeleteActor implements FilmRepository.
func (repo *FilmRepositoryImpl) DeleteActor(id uint) error {
	db := repo.DB.Delete(&model.FilmActor{}, id)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// GetActor implements FilmRepository.
func (repo *FilmRepositoryImpl) GetActor() []model.Actor {
	actor := []model.Actor{}
	repo.DB.Find(&actor)
	return actor
}
