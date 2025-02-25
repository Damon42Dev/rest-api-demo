package services

import (
	"context"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories/mongodb_repo"
	"example/rest-api-demo/src/utils"

	"go.mongodb.org/mongo-driver/mongo/options"
)

type MoviesService interface {
	GetMovies(pageStr, sizeStr string, ctx context.Context) ([]*models.Movie, error)
	GetMovieByID(id string, ctx context.Context) (*models.Movie, error)
}

type moviesService struct {
	mr mongodb_repo.MoviesRepository
}

func NewMoviesService(mr mongodb_repo.MoviesRepository) MoviesService {
	return &moviesService{mr: mr}
}

func (ms *moviesService) GetMovies(pageStr, sizeStr string, ctx context.Context) ([]*models.Movie, error) {
	pagination := utils.GetPaginationParams(pageStr, sizeStr)

	findOptions := options.Find()
	findOptions.SetLimit(int64(pagination.Size))
	findOptions.SetSkip(int64((pagination.Page - 1) * pagination.Size))

	return ms.mr.GetMovies(findOptions, ctx)
}

func (ms *moviesService) GetMovieByID(idStr string, ctx context.Context) (*models.Movie, error) {
	return ms.mr.GetMovieByID(idStr, ctx)
}
