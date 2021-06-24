package interfaces

import models "users_api/internal/core/domain/models/users"

type UserRepository interface {
	Get(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	All() (*[]models.User, error)
	Query(q *models.User) ([]models.User, error)
	Add(user *models.User) error

}