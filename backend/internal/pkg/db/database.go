package db

import (
	"fmt"
	"io/ioutil"
	"time"
	models "users_api/internal/core/domain/models/users"
	"users_api/internal/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

type Database struct {
	*gorm.DB
}

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() {
	var db = DB

	configuration := config.GetConfig()

	database := configuration.Database.Dbname
	username := configuration.Database.Username
	host := configuration.Database.Host
	port := configuration.Database.Port

	bin, err := ioutil.ReadFile("/run/secrets/db-password")
	if err != nil {
		fmt.Println("db err: ", err)
	}

	conn_string := username + ":" + string(bin) + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(conn_string), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: ", err)
	}

	// Change this to true if you want to see SQL queries
	db.Logger.LogMode(logger.Error)
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	sqlDb.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	DB = db
	migration()
}

// Auto migrate project models
func migration() {
	DB.AutoMigrate(&models.User{})

}

func GetDB() *gorm.DB {
	return DB
}
