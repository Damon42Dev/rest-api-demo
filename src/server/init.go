package server

import (
	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/mongodb"
	"example/rest-api-demo/src/repositories/mongodb_repo"
	"example/rest-api-demo/src/routes"
	"example/rest-api-demo/src/services"
	"example/rest-api-demo/src/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(config utils.Configuration) {
	client, err := mongodb.ConnectMongoDb(config.Database.Uri)

	if err != nil {
		log.Fatal(err)
	}

	// comments_repository := mongodb_repo.NewCommentMongodbRepo(&config, client)
	// comments_controller := controllers.NewCommentsController(client, comments_repository, config)

	moviesRepository := mongodb_repo.NewMovieMongodbRepo(&config, client)
	moviesService := services.NewMoviesService(moviesRepository)
	moviesController := controllers.NewMoviesController(client, moviesService, config)

	// Create an instance of the Controllers struct
	controllers := routes.Controllers{
		// CommentsController: commentsController,
		MoviesController: moviesController,
	}

	// Creates a gin router with default middleware:
	r := gin.Default()
	routes.RegisterRoutes(r, controllers)

	r.Run(":8080")
}
