package database

import (
	"fmt"
	// "log"
	"os"

	// "github.com/joho/godotenv"
	"github.com/rs401/TFG/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	dbpass := os.Getenv("POSTGRES_PASSWORD")
	dbuser := os.Getenv("POSTGRES_USER")
	dbport := os.Getenv("POSTGRES_PORT")
	// dbhost := os.Getenv("POSTGRES_HOST")
	dbhost := "localhost"
	dbname := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf("host=" + dbhost + " user=" + dbuser + " password=" + dbpass + " dbname=" + dbname + " port=" + dbport)
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect.")
	}
	fmt.Println("Connected to database.")

	DBConn.AutoMigrate(&models.Forum{}, &models.Thread{}, &models.Post{}, &models.User{})
	fmt.Println("Database migrated.")
}
