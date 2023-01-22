package service

import (
	"xyz_test/model"
	"xyz_test/repository"
)

type UserService interface {
	CreateUser(list model.UserCreateRequest) error
	LoginUser(request model.UserLoginRequest) (userData model.User, err error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) CreateUser(request model.UserCreateRequest) error {
	userCreate := model.User{
		Nik:          request.Nik,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: request.TanggalLahir,
		Gaji:         request.Gaji,
		Password:     request.Password,
		FotoKtp:      request.FotoKtp,
		FotoSelfie:   request.FotoSelfie,
	}
	if err := s.userRepository.CreateUser(userCreate); err != nil {
		return err
	}

	return nil
}

func (s *userService) LoginUser(request model.UserLoginRequest) (userData model.User, err error) {
	userLogin := model.User{
		Nik: request.Nik,
	}

	userData, err = s.userRepository.LoginUser(userLogin)
	if err != nil {
		return userData, err
	}

	return userData, nil
}
