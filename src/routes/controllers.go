package routes

import (
	"example/rest-api-demo/src/controllers"
)

// Controllers struct to hold all controllers
type Controllers struct {
	MoviesController   controllers.MoviesController
	CommentsController controllers.CommentsController
	// Add other controllers here
}
