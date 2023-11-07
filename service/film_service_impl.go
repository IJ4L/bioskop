package service

import (
	"errors"
	"fmt"

	"cinema.com/data/request"
	"cinema.com/data/response"
	"cinema.com/model"
	"cinema.com/repository"
	"github.com/go-playground/validator/v10"
)

type FilmServiceImpl struct {
	filmRepository repository.FilmRepository
	validate       *validator.Validate
}

func NewFilmService(filmRepositor repository.FilmRepository, validate *validator.Validate) *FilmServiceImpl {
	return &FilmServiceImpl{filmRepository: filmRepositor, validate: validate}
}

func (s *FilmServiceImpl) GetAllFilm() ([]model.Film, error) {
	data := s.filmRepository.GetAll()
	if data == nil {
		return nil, errors.New("no films found")
	}
	return data, nil
}

func (s *FilmServiceImpl) GetSeat(id int) ([]response.SeatStatus, error) {
	data := s.filmRepository.GetSeat(id)
	if data == nil {
		return nil, fmt.Errorf("no seats found for film with ID %d", id)
	}
	return data, nil
}

func (s *FilmServiceImpl) AddFilm(film request.AddFilm) error {
	err := s.validate.Struct(&film)
	if err != nil {
		return err
	}

	err = s.filmRepository.AddFilm(film)
	if err != nil {
		return err
	}

	return nil
}

func (s *FilmServiceImpl) DeleteFilm(id uint) error {

	err := s.filmRepository.FindById(id)
	if err != nil {
		return fmt.Errorf("film with ID %d not found", id)
	}

	err = s.filmRepository.DeleteFilm(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *FilmServiceImpl) BookingFilm(booking request.BookingFilm) (request.BookingFilm, error) {
	err := s.validate.Struct(&booking)
	if err != nil {
		return booking, fmt.Errorf("booking film failed: %w", err)
	}

	err = s.filmRepository.FindById(booking.IdFilm)
	if err != nil {
		return booking, fmt.Errorf("film with ID %d not found", booking.IdFilm)
	}

	_, err = s.filmRepository.BookingFilm(booking)
	if err != nil {
		return booking, fmt.Errorf("booking film failed: %w", err)
	}

	return booking, nil
}

// AddActor implements FilmService.
func (s *FilmServiceImpl) AddActor(actor request.AddActor) error {
	err := s.validate.Struct(&actor)
	if err != nil {
		return err
	}

	err = s.filmRepository.AddActor(actor)
	if err != nil {
		return err
	}

	return nil
}

// ConnectActor implements FilmService.
func (s *FilmServiceImpl) ConnectActor(connect request.ConnectActor) error {
	err := s.validate.Struct(&connect)
	if err != nil {
		return err
	}

	err = s.filmRepository.ConnectActor(connect)
	if err != nil {
		return err
	}

	return nil
}

// DeleteActor implements FilmService.
func (s *FilmServiceImpl) DeleteActor(id uint) error {
	db := s.filmRepository.DeleteActor(id)
	if db != nil {
		return db
	}

	return nil
}

// GetActor implements FilmService.
func (s *FilmServiceImpl) GetActor() []model.Actor {
	return s.filmRepository.GetActor()
}
