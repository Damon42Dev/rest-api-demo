package controllers

import (
	"example/rest-api-demo/models"
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

func DeleteCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = repositories.DeleteCommentByID(objID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting document: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func UpdateCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = repositories.UpdateCommentByID(objID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating document: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := repositories.CreateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("Error inserting document: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully",
		"id": id.Hex()})
}
