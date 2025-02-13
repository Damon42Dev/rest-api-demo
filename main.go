package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"example/rest-api-demo/src/server"
	"example/rest-api-demo/src/utils"

	"github.com/joho/godotenv"
)

// var db *mongo.Database

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := read_configuration()

	server.Initialize(config)
}

func read_configuration() utils.Configuration {
	mongoUrl := os.Getenv("MONGODB_URI")
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	collectionsStr := os.Getenv("COLLECTION")
	appName := os.Getenv("APP_NAME")
	requestTimeOutStr := os.Getenv("TIMEOUT")

	// Convert requestTimeOut to an integer
	requestTimeOut, err := strconv.Atoi(requestTimeOutStr)
	if err != nil {
		log.Fatalf("Invalid TIMEOUT value: %s", requestTimeOutStr)
	}

	// Split collectionsStr into a slice of strings
	collections := strings.Split(collectionsStr, ",")

	// log.Printf("Mongo URL: %s", mongoUrl)
	// log.Printf("Port: %s", port)
	// log.Printf("DB Name: %s", dbName)
	// log.Printf("Collection: %s", collectionsStr)
	// log.Printf("App Name: %s", appName)
	// log.Printf("Request Timeout: %d", requestTimeOut)

	return utils.Configuration{
		App:      utils.Application{Name: appName, Timeout: requestTimeOut},
		Database: utils.DatabaseSetting{Uri: mongoUrl, DbName: dbName, Collections: collections},
		Server:   utils.ServerSettings{Port: port},
	}
}
