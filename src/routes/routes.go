package routes

import (
	"example/rest-api-demo/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, commentController controllers.CommentsController, movieController controllers.MoviesController) {
	r.GET("/movies", movieController.GetMovies)
	r.GET("/movies/:id", movieController.GetMovieByID)

	r.GET("/comments", commentController.GetComments)
	r.GET("/comments/:id", commentController.GetCommentByID)
	r.DELETE("/comments/:id", commentController.DeleteCommentByID)
	r.PUT("/comments/:id", commentController.UpdateCommentByID)
	r.POST("/comments", commentController.CreateComment)
}
