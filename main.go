package main

import (
	"os"

	"example/rest-api-demo/src/server"
	"example/rest-api-demo/src/utils"
)

// var db *mongo.Database

func main() {
	config := read_configuration()

	server.Initialize(config)
}

func read_configuration() utils.Configuration {

	mongoUrl := os.Getenv("MONGODB_URI")
	port := os.Getenv("SERVER_PORT")
	dbName := os.Getenv("DB_NAME")
	collection := os.Getenv("COLLECTION")
	appName := os.Getenv("APP_NAME")
	// requestTimeOut := os.Getenv("TIMEOUT")

	return utils.Configuration{
		App:      utils.Application{Name: appName, Timeout: 3000},
		Database: utils.DatabaseSetting{Uri: mongoUrl, DbName: dbName, Collection: collection},
		Server:   utils.ServerSettings{Port: port},
	}
}
