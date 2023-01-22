package repository

import (
	"log"
	"xyz_test/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) error
	LoginUser(user model.User) (userData model.User, err error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user model.User) error {
	db := r.db
	err := db.Create(&user)
	if err.Error != nil {
		log.Printf("Error create data user with err: %v", err)
		return err.Error
	}
	return nil
}

func (r *userRepository) LoginUser(user model.User) (userData model.User, err error) {
	db := r.db
	var userCheck model.User
	if err := db.Where("nik = ?", user.Nik).First(&userCheck).Error; err != nil {
		log.Printf("Token not found with err: %s", err)
		return userData, err
	}
	return userCheck, nil
}
