package controllers

import (
	"errors"
	"log"
	"net/http"
	"users_api/internal/pkg/dtos"
	models "users_api/internal/pkg/models/users"
	"users_api/internal/pkg/persistence"
	"users_api/pkg/crypto"
	http_err "users_api/pkg/http-err"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary login User with username and password and get jwt token
// @Description login User
// @Accept json
// @Produce json
// @Param credentials body dtos.LoginInput true "Credentials"
// @Success 200 {object} string
// @Router /api/login [post]
func Login(c *gin.Context) {
	var loginInput dtos.LoginInput
	_ = c.BindJSON(&loginInput)
	s := persistence.GetUserRepository()
	if user, err := s.GetByUsername(loginInput.Username); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
			http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
			return
		}
		token, _ := crypto.CreateToken(user.Username)
		c.JSON(http.StatusOK, token)
	}
}

// GetUsers godoc
// @Summary register User based on username and password
// @Description Register User
// @Accept json
// @Produce json
// @Param user body dtos.UserInput true "User"
// @Success 201
// @Router /api/login/register [post]
func RegisterUser(c *gin.Context) {
	s := persistence.GetUserRepository()
	var userInput dtos.UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Username:  userInput.Username,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
	}
	if err := s.Add(&user); err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.Status(http.StatusCreated)
	}
}
