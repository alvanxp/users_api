package controllers

import (
	"errors"
	"log"
	"net/http"
	"users_api/internal/core/interfaces"
	http_err "users_api/pkg/http-err"

	"github.com/gin-gonic/gin"
)

// UserController represents the http handler for user
type UserController struct {
	userService interfaces.UserService
}

// NewUserController creates a new user controller
func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{service}
}

// GetUserById godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security ApiKeyAuth
func (u *UserController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	if user, err := u.userService.Get(id); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// GetUsers godoc
// @Summary Retrieves users based on query
// @Description Get Users
// @Produce json
// @Param username query string false "Username"
// @Param firstname query string false "Firstname"
// @Param lastname query string false "Lastname"
// @Success 200 {array} []users.User
// @Router /api/users [get]
// @Security ApiKeyAuth
func (u *UserController) GetUsers(c *gin.Context) {
	if users, err := u.userService.GetAll(); err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
	// }
	// if users, err := u.userService.GetUsers(q.Username, q.Firstname, q.Lastname); err != nil {
	// 	http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
	// 	log.Println(err)
	// } else {
	// 	c.JSON(http.StatusOK, users)
	//}
}
