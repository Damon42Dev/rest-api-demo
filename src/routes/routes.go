package routes

import (
	"example/rest-api-demo/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, commentController controllers.CommentsController) {
	// r.GET("/movies", controllers.GetMovies)
	// r.GET("/movies/:id", controllers.GetMovieByID)

	r.GET("/comments", commentController.GetComments)
	r.GET("/comments/:id", commentController.GetCommentByID)
	r.DELETE("/comments/:id", commentController.DeleteCommentByID)
	// r.POST("/comments", controllers.CreateComment)
	// r.PUT("/comments/:id", controllers.UpdateCommentByID)
}
