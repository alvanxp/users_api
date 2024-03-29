package interfaces

import (
	"users_api/internal/core/domain/dtos"
)

// AuthService represents the interface for authentication service.
type AuthService interface {
	ValidateUser(login dtos.LoginInput) (string, error)
	RegisterUser(userInput dtos.UserInput) (string, error)
}
