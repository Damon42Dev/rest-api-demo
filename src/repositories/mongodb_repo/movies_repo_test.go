package mongodb_repo_test

import (
	"context"
	"testing"

	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/repositories/mongodb_repo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	testDB     *mongodb_repo.TestDBClient
	moviesRepo repositories.MoviesRepository
)

func TestGetMovies(t *testing.T) {
	// Setup
	testDB = mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	moviesRepo = mongodb_repo.NewMovieMongodbRepo(testDB.Config, testDB.Client)

	// Test GetMovies
	findOptions := options.Find()
	result, err := moviesRepo.GetMovies(findOptions, context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Test Movie 1", result[0].Title)
	assert.Equal(t, "Test Movie 2", result[1].Title)
}

func TestGetMovieByID(t *testing.T) {
	// Setup
	testDB = mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	moviesRepo = mongodb_repo.NewMovieMongodbRepo(testDB.Config, testDB.Client)

	// Test GetMovieByID
	result, err := moviesRepo.GetMovieByID("60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Test Movie 1", result.Title)
}
