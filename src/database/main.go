package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/nuck7/nfl-picks-server/src/models"
)

func Init() (db *gorm.DB) {
	conf :=
		Config{
			// FILL IN
		}

	connectionString := GetConnectionString(conf)

	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Team{}, &models.Pick{}, &models.Matchup{}, &models.User{}, &models.Week{})

	return db
}

func main() {
	// config :=
	// 	Config{
	// 		ServerName: "localhost:3308",
	// 		User:       "root",
	// 		Password:   "SecuresAF123",
	// 		DB:         "nfl-picks",
	// 	}

	// db, err := gorm.Open(sqlite.Open("mysql"), config)

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// // Migrate the schema
	// db.AutoMigrate(&models.Team{}, &models.Pick{}, &models.Matchup{}, &models.User{}, &models.Week{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	// var product Product
	// db.First(&product, 1)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}
