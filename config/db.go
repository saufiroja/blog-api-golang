package config

import (
	"echo/blog-api/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database connection configuration.
type Config struct {
	DB_HOST    string
	DB_PORT    string
	DB_USER    string
	DB_PASS    string
	DB_NAME    string
	JWT_SECRET string
}

// INITIALIZE is the database connection configuration.
func InitDB(conf Config) *gorm.DB {
	// load godotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// connect to database
	// dns
	dns := "host=" + host + " user=" + username + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	db.AutoMigrate(&entity.User{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
