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
	if user, err:=a.userRepository.GetByUsername(login.Username); err != nil{
		return "", err
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(login.Password)){
			return "", errors.New("user not authenticated")
		}
		return crypto.CreateToken(user.Username)
	}
}

func (a *authServiceImp) RegisterUser(userInput dtos.UserInput) error {
	user := &models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
	}
	return a.userRepository.Add(user)
}

func NewAuthServiceImp(userRepository interfaces.UserRepository) *authServiceImp {
	return &authServiceImp{ userRepository: userRepository}
}
