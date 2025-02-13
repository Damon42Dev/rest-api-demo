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

	// Creates a gin router with default middleware:
	r := gin.Default()
	routes.RegisterRoutes(r, comments_controller)

	r.Run(":8080")
}
