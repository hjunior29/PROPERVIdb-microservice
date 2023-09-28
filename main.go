package main

import (
	"log"
	"os"

	"github.com/hjunior29/PROPERVIdb-microservice/db"
	"github.com/hjunior29/PROPERVIdb-microservice/routes"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load("config/.env"); err != nil {
		log.Println("No .env file found")
	}
}

func main() {

	if err := db.InitDatabase(); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	log.Println("Connected to the database!")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	routes.HandleRequest().Run()

}
