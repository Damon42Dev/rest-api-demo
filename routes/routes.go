package routes

import (
	"example/rest-api-demo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/movies", controllers.GetMovies)
	// server.GET("/comments", GetComments)
	// server.GET("/comments/:id", GetCommentByID)
	// server.POST("/comments", CreateComment)
	// server.PUT("/comments/:id", UpdateCommentByID)
	// server.DELETE("/comments/:id", DeleteCommentByID)

	// server.GET("/movies/:id", GetMovieByID)
}
