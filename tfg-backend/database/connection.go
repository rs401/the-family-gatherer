package database

import (
	"fmt"

	"github.com/rs401/TFG/tfg-backend/config"
	"github.com/rs401/TFG/tfg-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dbpass := config.Config("POSTGRES_PASSWORD")
	dbuser := config.Config("POSTGRES_USER")
	dbport := config.Config("POSTGRES_PORT")
	// dbhost := config.Config("POSTGRES_HOST")
	dbhost := "localhost"
	dbname := config.Config("POSTGRES_DB")

	dsn := fmt.Sprintf("host=" + dbhost + " user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " port=" + dbport)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect.")
	}
	fmt.Println("Connected to database.")

	DBConn.AutoMigrate(&models.User{}) //, &models.Forum{}, &models.Thread{}, &models.Post{})
	fmt.Println("Database migrated.")
}
