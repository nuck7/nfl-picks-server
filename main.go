package main

import (
	"log"

	"github.com/nuck7/nfl-picks-server/src/api"
	"github.com/nuck7/nfl-picks-server/src/database"

	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	// connectionString := database.GetConnectionString(config)
	// err := database.Connect(connectionString)
	// if err != nil {
	// 	panic(err.Error())
	// }

	db := database.Init()

	log.Println("Connection was successful!!", db)

	api.Api()

}
