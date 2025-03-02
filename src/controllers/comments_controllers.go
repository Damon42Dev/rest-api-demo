package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/services"
	"example/rest-api-demo/src/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentsController interface {
	GetComments(*gin.Context)
	GetCommentByID(*gin.Context)
	DeleteCommentByID(*gin.Context)
	UpdateCommentByID(*gin.Context)
	CreateComment(*gin.Context)
}

type commentsController struct {
	client          *mongo.Client
	commentsService services.CommentsService
	config          utils.Configuration
}

func NewCommentsController(client *mongo.Client, service services.CommentsService, config utils.Configuration) CommentsController {
	return &commentsController{client: client, commentsService: service, config: config}
}

func (cc *commentsController) GetComments(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer cancel()

	pageStr := c.Query("page")
	sizeStr := c.Query("size")

	comments, err := cc.commentsService.GetComments(pageStr, sizeStr, ctx)

	if err != nil {
		log.Printf("Error getting comments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *commentsController) GetCommentByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer cancel()

	idStr := utils.GetIdStrFromParam(c, "id")

	comment, err := cc.commentsService.GetCommentByID(idStr, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to retrieve comment by ID: %s", idStr)})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (cc *commentsController) DeleteCommentByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer cancel()

	idStr := utils.GetIdStrFromParam(c, "id")

	err := cc.commentsService.DeleteCommentByID(idStr, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting comment by ID: %s", idStr)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func (cc *commentsController) UpdateCommentByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer cancel()

	idStr := utils.GetIdStrFromParam(c, "id")

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := cc.commentsService.UpdateCommentByID(idStr, updateData, ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating document: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (cc *commentsController) CreateComment(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(cc.config.App.Timeout)*time.Second)
	defer cancel()

	var comment models.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, err := cc.commentsService.CreateComment(comment, ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{"error": fmt.Sprintf("Error inserting document: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully",
		"id": id})
}
