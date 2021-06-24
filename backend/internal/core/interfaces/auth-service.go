package interfaces

import (
	"users_api/internal/core/domain/dtos"
)

type AuthService interface {
		ValidateUser(login dtos.LoginInput) (string, error)
		RegisterUser(userInput dtos.UserInput) error
}
