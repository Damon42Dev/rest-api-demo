package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, controllers Controllers) {
	r.GET("/movies", controllers.MoviesController.GetMovies)
	r.GET("/movies/:id", controllers.MoviesController.GetMovieByID)

	r.GET("/comments", controllers.CommentsController.GetComments)
	r.GET("/comments/:id", controllers.CommentsController.GetCommentByID)
	r.DELETE("/comments/:id", controllers.CommentsController.DeleteCommentByID)
	r.PUT("/comments/:id", controllers.CommentsController.UpdateCommentByID)
	// r.POST("/comments", controllers.CommentsController.CreateComment)
}
