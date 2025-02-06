package controllers

import (
	"example/rest-api-demo/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
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
