package service

import "cinema.com/data/request"

type AuthenticationService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.CreateUsersRequest) (string, error)
}
