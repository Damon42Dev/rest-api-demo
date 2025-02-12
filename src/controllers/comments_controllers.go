package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentsController interface {
	// Healthcheck(*gin.Context)

	// Add(*gin.Context)
	GetComments(*gin.Context)
	// GetById(*gin.Context)
	// Delete(*gin.Context)
}

type commentsController struct {
	client             *mongo.Client
	commentsRepository repositories.CommentsRepository
	config             utils.Configuration
}

func NewCommentsController(client *mongo.Client, repo repositories.CommentsRepository, config utils.Configuration) CommentsController {
	return &commentsController{client: client, commentsRepository: repo, config: config}
}

func (cc *commentsController) GetComments(c *gin.Context) {

	ctx, ctxErr := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer ctxErr()

	var commentModel []*models.Comment

	tParam := c.Param("take")
	take, err := strconv.Atoi(tParam)
	if err != nil {
		log.Fatal("Take parameter can not be converted.")
	}

	result, err := cc.commentsRepository.GetComments(take, ctx)
	if err != mongo.ErrNilCursor {
		// utils.BadRequestError("AppDoc_Handler_List", err, map[string]interface{}{"Data": take})
		log.Fatal("Error getting comments")
	}
	// logrus.Infof("Len %d", len(result))

	//convert to entity to model
	for _, item := range result {
		commentModel = append(commentModel, (*models.Comment)(item))
	}

	c.IndentedJSON(http.StatusOK, map[string]interface{}{"Data": commentModel})
}

// func GetCommentByID(c *gin.Context) {
// 	id := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	movie, err := repositories.GetCommentByID(objID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve comment by ID: %s", objID.Hex())})
// 		return
// 	}

// 	c.JSON(http.StatusOK, movie)
// }

// func DeleteCommentByID(c *gin.Context) {
// 	id := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	err = repositories.DeleteCommentByID(objID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting document: %s", err)})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
// }

// func UpdateCommentByID(c *gin.Context) {
// 	id := c.Param("id")
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	var updateData map[string]interface{}
// 	if err := c.BindJSON(&updateData); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	err = repositories.UpdateCommentByID(objID, updateData)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating document: %s", err)})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
// }

// func CreateComment(c *gin.Context) {
// 	var comment models.Comment
// 	if err := c.BindJSON(&comment); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
// 		return
// 	}

// 	id, err := repositories.CreateComment(comment)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError,
// 			gin.H{"error": fmt.Sprintf("Error inserting document: %s", err)})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully",
// 		"id": id.Hex()})
// }
