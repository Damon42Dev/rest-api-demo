package routes

import (
	"example/rest-api-demo/src/controllers"
)

type Controllers struct {
	CommentsController controllers.CommentsController
	MoviesController   controllers.MoviesController
}
