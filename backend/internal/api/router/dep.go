package router

import (
	"users_api/internal/api/controllers"
	"users_api/internal/core/interfaces"
	"users_api/internal/core/services"
	"users_api/internal/pkg/persistence"
)

type dep struct {
	UserRepository interfaces.UserRepository
	UserService interfaces.UserService
	UserController *controllers.UserController
	AuthController *controllers.AuthController
}

func NewDependencies() *dep {
	userRepository :=&persistence.UserRepositoryImp{}
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthServiceImp(userRepository)
	return &dep{
		UserRepository: userRepository,
		UserService:    userService,
		UserController: controllers.NewUserController(userService),
		AuthController: controllers.NewAuthController(authService),
	}
}