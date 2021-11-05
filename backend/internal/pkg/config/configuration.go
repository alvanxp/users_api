package config

import (
	"log"

	"github.com/spf13/viper"
)

//Config instance of Configuration file
var Config *Configuration

//Configuration class to map config file
type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

//DatabaseConfiguration class to map db configuration
type DatabaseConfiguration struct {
	Dbname       string
	Username     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

//ServerConfiguration class to map server configuration
type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

//Setup initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
