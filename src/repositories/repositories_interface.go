package repositories

import (
	"context"
	"example/rest-api-demo/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentsRepository interface {
	GetComments(findOptions *options.FindOptions, ctx context.Context) ([]*models.Comment, error)
	GetCommentByID(id string, ctx context.Context) (*models.Comment, error)
	DeleteCommentByID(id string, ctx context.Context) error
	UpdateCommentByID(id string, updateData bson.M, ctx context.Context) error
	CreateComment(comment models.Comment, ctx context.Context) (string, error)
	GetCommentsForMovie(findOptions *options.FindOptions, idStr string, ctx context.Context) ([]*models.Comment, error)
}

type MoviesRepository interface {
	GetMovies(findOptions *options.FindOptions, ctx context.Context) ([]*models.Movie, error)
	GetMovieByID(idStr string, ctx context.Context) (*models.Movie, error)
}
