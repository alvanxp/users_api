package services

import (
	"errors"
	"users_api/internal/core/domain/dtos"
	models "users_api/internal/core/domain/models/users"
	"users_api/internal/core/interfaces"
	"users_api/pkg/crypto"
)

type authServiceImp struct {
	userRepository interfaces.UserRepository
}

func (a *authServiceImp) ValidateUser(login dtos.LoginInput) (string, error) {
	user, err := a.userRepository.GetByUsername(login.Username)
	if err != nil {
		return "", err
	}
	if !crypto.ComparePasswords(user.Hash, []byte(login.Password)) {
		return "", errors.New("user not authenticated")
	}
	token, err := crypto.CreateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authServiceImp) RegisterUser(userInput dtos.UserInput) (string, error) {
	user := &models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
	}
	//verify if user exists
	if _, err := a.userRepository.GetByUsername(user.Username); err == nil {
		return "", errors.New("user already exists")
	}
	token, err := crypto.CreateToken(user.Username)
	if err != nil {
		return "", err
	}
	err = a.userRepository.Add(user)
	if err != nil {
		return "", err
	}
	return token, nil
}

// NewAuthServiceImp creates a new instance of authServiceImp.
func NewAuthServiceImp(userRepository interfaces.UserRepository) interfaces.AuthService {
	return &authServiceImp{userRepository: userRepository}
}
