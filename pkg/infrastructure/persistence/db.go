package persistence

import (
	"example/gin-project/pkg/domain"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Postgres_Init() *gorm.DB {
	err := godotenv.Load(".env")

	log.Println("DBURL: ", os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatalln("Error loading .env file", err.Error())
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error connecting to database", err.Error())
	}

	db.AutoMigrate(&domain.User{})

	return db
}
