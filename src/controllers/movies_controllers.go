package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"example/rest-api-demo/src/services"
	"example/rest-api-demo/src/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type MoviesController interface {
	GetMovies(*gin.Context)
	GetMovieByID(*gin.Context)
}

type moviesController struct {
	client        *mongo.Client
	moviesService services.MoviesService
	config        utils.Configuration
}

func NewMoviesController(client *mongo.Client, service services.MoviesService, config utils.Configuration) MoviesController {
	return &moviesController{client: client, moviesService: service, config: config}
}

func (mc *moviesController) GetMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(mc.config.App.Timeout)*time.Second)
	defer cancel()

	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	movies, err := mc.moviesService.GetMovies(pageStr, sizeStr, ctx)
	if err != nil {
		log.Printf("Error getting movies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (mc *moviesController) GetMovieByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(mc.config.App.Timeout)*time.Second)
	defer cancel()

	idStr := utils.GetIdStrFromParam(c, "id")

	movie, err := mc.moviesService.GetMovieByID(idStr, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve movie by ID: %s", idStr)})
		return
	}

	c.JSON(http.StatusOK, movie)
}
