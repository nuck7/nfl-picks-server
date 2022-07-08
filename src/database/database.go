package database

import (
	"fmt"
)

//Config to maintain DB configuration properties
type Config struct {
	Host     string
	User     string
	Password string
	DB       string
}

var GetConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.DB)
	return connectionString
}
