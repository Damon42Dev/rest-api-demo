package routes

import (
	"example/rest-api-demo/src/controllers"
)

type Controllers struct {
	CommentsController controllers.CommentsController
	MoviesController   controllers.MoviesController
}

func NewControllers(commentsController controllers.CommentsController, moviesController controllers.MoviesController) Controllers {
	return Controllers{
		CommentsController: commentsController,
		MoviesController:   moviesController,
	}
}
