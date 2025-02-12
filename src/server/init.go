package server

import (
	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/mongodb"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func Initialize(config utils.Configuration) {

	log.Println("Starting server...", config.Database.Uri)
	client, err := mongodb.ConnectMongoDb(config.Database.Uri)

	log.Println("Connected to MongoDB", client)
	if err != nil {
		log.Fatal(err)
	}

	comments_repository := repositories.NewCommentMongodbRepo(&config, client)
	comments_controller := controllers.NewCommentsController(client, comments_repository, config)

	// Creates a gin router with default middleware:
	router := gin.Default()

	// Register API routes
	api := router.Group("api/v1")
	{
		// api.GET("/health", handler.Healthcheck)

		// api.POST("/appdoc/add", handler.Add)
		api.GET("/rest-api-demo/comments-list/:take", comments_controller.GetComments)
		// api.GET("/appdoc/get/:id", handler.GetById)
		// api.PUT("/appdoc/delete/:id", handler.Delete)
	}

	// PORT environment variable was defined.
	formattedUrl := fmt.Sprintf(":%s", config.Server.Port)

	router.Run(formattedUrl)
}
