package tests

import (
	"os"
	"testing"

	"github.com/hjunior29/PROPERVIdb-microservice/db"
	"github.com/hjunior29/PROPERVIdb-microservice/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func loadEnv() {
	if err := godotenv.Load("../config/.env.test"); err != nil {
		panic("No .env.test file found")
	}
}

func TestInitDatabase(t *testing.T) {

	loadEnv()

	err := db.InitDatabase()

	assert.Nil(t, err)
	assert.NotNil(t, db.DB)
}

func TestInitDatabaseWithMissingEnv(t *testing.T) {

	loadEnv()

	originDB_USER := os.Getenv("DB_USER")
	os.Unsetenv("DB_USER")

	err := db.InitDatabase()
	assert.NotNil(t, err)

	os.Setenv("DB_USER", originDB_USER)
}

func TestDatabaseOperations(t *testing.T) {
	loadEnv()

	err := db.InitDatabase()
	assert.Nil(t, err)

	testProperty := &models.Properties{
		Address: "New York, 456",
		Price:   50000000,
	}

	err = db.DB.Create(testProperty).Error
	assert.Nil(t, err, "Failed to add properties to the database")

	var retrievedProperty models.Properties
	err = db.DB.First(&retrievedProperty, testProperty.ID).Error
	assert.Nil(t, err, "Failed to find properties in the database")

	assert.Equal(t, testProperty.Address, retrievedProperty.Address, "Address fields do not match")
	assert.Equal(t, testProperty.Price, retrievedProperty.Price, "Price fields do not match")

	err = db.DB.Delete(testProperty).Error
	assert.Nil(t, err, "Failed to delete properties in the database")
}
