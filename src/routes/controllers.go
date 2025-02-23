package routes

import (
	"example/rest-api-demo/src/controllers"
)

// Controllers struct to hold all controllers
type Controllers struct {
	MoviesController controllers.MoviesController
	// Add other controllers here
}
