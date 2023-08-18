package service

import (
	"errors"
	"time"

	"cinema.com/data/request"
	"cinema.com/helper"
	"cinema.com/model"
	"cinema.com/repository"
	"cinema.com/utils"
	"github.com/go-playground/validator/v10"
)

type AuthenticationServiceImpl struct {
	UserRepository repository.UserRepository
	valdate        *validator.Validate
}

func NewAuthenticationService(userRepository repository.UserRepository, valdate *validator.Validate) *AuthenticationServiceImpl {
	return &AuthenticationServiceImpl{UserRepository: userRepository, valdate: valdate}
}

func (s *AuthenticationServiceImpl) Login(user request.LoginRequest) (string, error) {

	new_users, user_er := s.UserRepository.FindByEmail(user.Email)
	if user_er != nil {
		return "", errors.New("invalid username or password")
	}

	verify_error := utils.VerifyPassword(new_users.Password, user.Password)
	if verify_error != nil {
		return "", errors.New("invalid username or password")
	}

	token, err_token := utils.GenerateToken(2*time.Hour, new_users.Id, "secret")
	helper.ErrorPanic(err_token)

	return token, nil
}

func (s *AuthenticationServiceImpl) Register(user request.CreateUsersRequest) (string, error) {
	err := s.valdate.Struct(user)
	if err != nil {
		return "", err
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	helper.ErrorPanic(err)

	newUser := model.User{
		Name:         user.Username,
		Email:        user.Email,
		Password:     hashedPassword,
		PhotoProfile: "https://images.unsplash.com/photo-1575936123452-b67c3203c357?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8aW1hZ2V8ZW58MHx8MHx8fDA%3D&w=1000&q=80",
		CreatedAt:    time.Now().String(),
	}

	s.UserRepository.Save(newUser)

	return "", nil
}
