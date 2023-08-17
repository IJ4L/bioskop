package repository

import "cinema.com/model"

type UserRepository interface {
	Save(users model.User)
	Update(users model.User)
	Delete(userId int)
	FindById(userId int) (model.User, error)
	FindAll() []model.User
	FindByEmail(username string) (model.User, error)
}
