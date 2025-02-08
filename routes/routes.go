package routes

import (
	"example/rest-api-demo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovieByID)

	r.GET("/comments", controllers.GetComments)
	r.GET("/comments/:id", controllers.GetCommentByID)
	r.DELETE("/comments/:id", controllers.DeleteCommentByID)
	r.POST("/comments", controllers.CreateComment)
	r.PUT("/comments/:id", controllers.UpdateCommentByID)
}
