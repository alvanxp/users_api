package main

import (
	_ "users_api/docs"
	"users_api/internal/api"
)

// @Golang API REST
// @version 1.0
// @description API REST in Golang with Gin Framework

// @contact.name Alvaro Carpio Paredes
// @contact.url
// @contact.email alvanxp@outlook.com

// @license.name MIT
// @license.url https://users_api/blob/master/LICENSE

// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	api.Run("")
}
