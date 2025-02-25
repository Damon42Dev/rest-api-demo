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

var config utils.Configuration

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config = readConfiguration()
}

func main() {
	server.Initialize(config)
}

func readConfiguration() utils.Configuration {
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

	return utils.Configuration{
		App:      utils.Application{Name: appName, Timeout: requestTimeOut},
		Database: utils.DatabaseSetting{Uri: mongoUrl, DbName: dbName, Collections: collections},
		Server:   utils.ServerSettings{Port: port},
	}
}
