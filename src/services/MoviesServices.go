package services

import (
	"context"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories/mongodb_repo"
)

type MoviesService interface {
	GetMovies(page, size int, ctx context.Context) ([]*models.Movie, error)
	// GetMovieByID(id string, ctx context.Context) (*models.Movie, error)
}

type moviesService struct {
	mr mongodb_repo.MoviesRepository
}

func NewMoviesService(mr mongodb_repo.MoviesRepository) MoviesService {
	return &moviesService{mr: mr}
}

func (ms *moviesService) GetMovies(page, size int, ctx context.Context) ([]*models.Movie, error) {
	if page < 1 {
		page = 1
	}

	if size < 1 {
		size = 10
	}

	return ms.mr.GetMovies(page, size, ctx)
}

// func (ms *moviesService) GetMovieByID(id string, ctx context.Context) (*models.Movie, error) {
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ms.mr.GetMovieByID(objectID, ctx)
// }
