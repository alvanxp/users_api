package persistence

import (
	"strconv"
	models "users_api/internal/core/domain/models/users"
)

type UserRepositoryImp struct{}

//
//var userRepository *UserRepositoryImp
//
//func GetUserRepository() *UserRepositoryImp {
//	if userRepository == nil {
//		userRepository = &UserRepositoryImp{}
//	}
//	return userRepository
//}

func (r *UserRepositoryImp) Get(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepositoryImp) GetByUsername(username string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Username = username
	_, err := First(&where, &user, []string{})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepositoryImp) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{}, "id asc")
	return &users, err
}

func (r *UserRepositoryImp) Query(q *models.User) ([]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{}, "id asc")
	return users, err
}

func (r *UserRepositoryImp) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}
