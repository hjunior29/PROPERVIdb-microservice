package db

import (
	"fmt"
	"os"

	"github.com/hjunior29/PROPERVIdb-microservice/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() error {

	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_SSLMODE"),
	)

	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		return err
	}

	DB.AutoMigrate(&models.Properties{})

	return nil
}
