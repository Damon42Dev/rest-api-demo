package controllers

import (
	"context"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/utils"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MoviesController interface {
	// Healthcheck(*gin.Context)

	// CreateComment(*gin.Context)
	GetMovies(*gin.Context)
	GetMovieByID(*gin.Context)
	// DeleteCommentByID(*gin.Context)
	// UpdateCommentByID(*gin.Context)
}

type moviesController struct {
	client           *mongo.Client
	moviesRepository repositories.MoviesRepository
	config           utils.Configuration
}

func NewMoviesController(client *mongo.Client, repo repositories.MoviesRepository, config utils.Configuration) MoviesController {
	return &moviesController{client: client, moviesRepository: repo, config: config}
}

func (mc *moviesController) GetMovies(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(mc.config.App.Timeout)*time.Second)
	defer ctxErr()

	var movieModel []*models.Movie

	pagination := utils.GetPaginationParams(c, 1, 5)

	result, err := mc.moviesRepository.GetMovies(pagination.Page, pagination.Size, ctx)
	if err != mongo.ErrNilCursor {
		log.Printf("Error getting movies")
	}

	//convert to entity to model
	for _, item := range result {
		movieModel = append(movieModel, (*models.Movie)(item))
	}

	c.IndentedJSON(http.StatusOK, map[string]interface{}{"Data": movieModel})

}

func (mc *moviesController) GetMovieByID(c *gin.Context) {
	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(mc.config.App.Timeout)*time.Second)
	defer ctxErr()

	objID, valid := utils.GetObjectIDFromParam(c, "id")
	if !valid {
		return
	}

	movie, err := mc.moviesRepository.GetMovieByID(objID, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve movie by ID: %s", objID.Hex())})
		return
	}

	c.JSON(http.StatusOK, movie)
}
