package interfaces

import models "users_api/internal/core/domain/models/users"

type UserService interface {
	Get(id string) (*models.User, error)
	GetUsers(username string, name string, lastname string) ([]models.User, error)

}
