package services

import (
	models "users_api/internal/core/domain/models/users"
	"users_api/internal/core/interfaces"
)

type userServiceImp struct {
	userRepository interfaces.UserRepository
}

func (u userServiceImp) Get(id string) (*models.User, error) {
	user, err := u.userRepository.Get(id)
	return user, err
}

func (u userServiceImp) GetUsers(username string, name string, lastname string) ([]models.User, error) {
	users, err := u.userRepository.Query(&models.User{Username: username, Firstname: name, Lastname: lastname})

	return users, err
}

// get all users
func (u userServiceImp) GetAll() (*[]models.User, error) {
	users, err := u.userRepository.All()
	return users, err
}


func NewUserService(repository interfaces.UserRepository) *userServiceImp {
	return &userServiceImp{
		userRepository: repository,
	}
}
