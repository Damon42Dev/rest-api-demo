package server

import (
	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/mongodb"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/routes"
	"example/rest-api-demo/src/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(config utils.Configuration) {
	client, err := mongodb.ConnectMongoDb(config.Database.Uri)

	if err != nil {
		log.Fatal(err)
	}

	comments_repository := repositories.NewCommentMongodbRepo(&config, client)
	comments_controller := controllers.NewCommentsController(client, comments_repository, config)

	movies_repository := repositories.NewMovieMongodbRepo(&config, client)
	movies_controller := controllers.NewMoviesController(client, movies_repository, config)

	// Create an instance of the Controllers struct
	controllers := routes.NewControllers(comments_controller, movies_controller)

	// Creates a gin router with default middleware:
	r := gin.Default()
	routes.RegisterRoutes(r, controllers)

	r.Run(":8080")
}
