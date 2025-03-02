package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/mongodb"
	"example/rest-api-demo/src/repositories/mongodb_repo"
	"example/rest-api-demo/src/routes"
	"example/rest-api-demo/src/services"
	"example/rest-api-demo/src/utils"

	"github.com/gin-gonic/gin"
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
	initialize(config)
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

func initialize(config utils.Configuration) {
	client, err := mongodb.ConnectMongoDb(config.Database.Uri)

	if err != nil {
		log.Fatal(err)
	}

	commentsMongoRepository := mongodb_repo.NewCommentMongodbRepo(&config, client)
	commentsService := services.NewCommentsService(commentsMongoRepository)
	commentsController := controllers.NewCommentsController(client, commentsService, config)

	moviesMongoRepository := mongodb_repo.NewMovieMongodbRepo(&config, client)
	moviesService := services.NewMoviesService(moviesMongoRepository)
	moviesController := controllers.NewMoviesController(client, moviesService, config)

	// Create an instance of the Controllers struct
	controllers := routes.Controllers{
		CommentsController: commentsController,
		MoviesController:   moviesController,
	}

	// Creates a gin router with default middleware:
	r := gin.Default()
	routes.RegisterRoutes(r, controllers)

	r.Run(":8080")
}
