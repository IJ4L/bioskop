package repository

import (
	"errors"

	"cinema.com/data/request"
	"cinema.com/helper"
	"cinema.com/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUsersRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Save(users model.User) {
	result := r.DB.Create(&users)
	helper.ErrorPanic(result.Error)
}

func (r *UserRepositoryImpl) Update(users model.User) {
	updateUsers := request.UpdateUsersRequest{
		Username: users.Name,
		Email:    users.Email,
		Password: users.Password,
	}

	result := r.DB.Model(&users).Updates(updateUsers)
	helper.ErrorPanic(result.Error)
}

func (r *UserRepositoryImpl) Delete(userId int) {
	users := model.User{}

	result := r.DB.First(&users, userId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

func (r *UserRepositoryImpl) FindById(userId int) (model.User, error) {
	users := model.User{}

	result := r.DB.Find(&users, userId)
	if result != nil {
		return users, result.Error
	} else {
		return users, errors.New("user not found")
	}
}

func (r *UserRepositoryImpl) FindAll() []model.User {
	users := []model.User{}

	result := r.DB.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

func (r *UserRepositoryImpl) FindByEmail(email string) (model.User, error) {
    user := model.User{}
    result := r.DB.Find(&user, "email = ?", email)

    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return user, errors.New("user not found")
        }
        return user, result.Error
    }

    return user, nil
}

