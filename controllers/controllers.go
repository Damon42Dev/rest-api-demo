package controllers

import (
	"example/rest-api-demo/repositories"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMovies(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	movies, err := repositories.GetMovies(pageStr, sizeStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	movie, err := repositories.GetMovieByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve movie by ID: %s", objID.Hex())})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func GetComments(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	comments, err := repositories.GetComments(pageStr, sizeStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	movie, err := repositories.GetCommentByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve comment by ID: %s", objID.Hex())})
		return
	}

	c.JSON(http.StatusOK, movie)
}
