package service

import (
	"cinema.com/data/response"
	"cinema.com/model"
	"cinema.com/repository"
	"errors"
	"fmt"
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
