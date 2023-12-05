package databases

import (
	"fmt"
	"jwt-go/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
  DB gorm.DB
}

var DB Database

func ConnectDatabase() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading env variables")
    return
  }

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DBNAME"), os.Getenv("POSTGRES_PORT"),
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    fmt.Println("error when connecting with database")
    return
  }

  err = db.AutoMigrate(models.User{})
  if err != nil {
    fmt.Println("error migrating user")
    return
  }

  DB = Database {
    DB: *db,
  }
  fmt.Println("Database connection established")
}
