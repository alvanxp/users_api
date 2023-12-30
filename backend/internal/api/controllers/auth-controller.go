package controllers

import (
	"errors"
	"log"
	"net/http"
	"users_api/internal/core/domain/dtos"
	"users_api/internal/core/interfaces"
	http_err "users_api/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService interfaces.AuthService
}

func NewAuthController(authService interfaces.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// Login godoc
// @Summary login User with username and password and get jwt token
// @Description login User
// @Accept json
// @Produce json
// @Param credentials body dtos.LoginInput true "Credentials"
// @Success 200 {object} string
// @Router /api/login [post]
func (a *AuthController) Login(c *gin.Context) {
	var loginInput dtos.LoginInput
	_ = c.BindJSON(&loginInput)
	if token, err := a.authService.ValidateUser(loginInput); err != nil {
		http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
		return
	} else {
		c.JSON(http.StatusOK, token)
	}
}

// RegisterUser godoc
// @Summary register User based on username and password
// @Description Register User
// @Accept json
// @Produce json
// @Param user body dtos.UserInput true "User"
// @Success 201
// @Router /api/login/register [post]
func (a *AuthController) RegisterUser(c *gin.Context) {
	var userInput dtos.UserInput
	_ = c.BindJSON(&userInput)
	token, err := a.authService.RegisterUser(userInput)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, token)
	}

}
