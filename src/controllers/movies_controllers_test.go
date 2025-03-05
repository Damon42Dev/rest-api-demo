package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockMoviesService is a mock implementation of the MoviesService interface
type MockMoviesService struct {
	mock.Mock
}

func (m *MockMoviesService) GetMovies(pageStr, sizeStr string, ctx context.Context) ([]*models.Movie, error) {
	args := m.Called(pageStr, sizeStr, ctx)
	return args.Get(0).([]*models.Movie), args.Error(1)
}

func (m *MockMoviesService) GetMovieByID(id string, ctx context.Context) (*models.Movie, error) {
	args := m.Called(id, ctx)
	return args.Get(0).(*models.Movie), args.Error(1)
}

func setupMoviesControllerTest() (*MockMoviesService, *gin.Engine) {
	mockService := new(MockMoviesService)
	config := utils.Configuration{
		App: utils.Application{
			Timeout: 5,
		},
	}
	controller := NewMoviesController(nil, mockService, config)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/movies", controller.GetMovies)
	router.GET("/movies/:id", controller.GetMovieByID)

	return mockService, router
}

func TestGetMovies(t *testing.T) {
	mockService, router := setupMoviesControllerTest()

	pageStr := "1"
	sizeStr := "5"
	movies := GetMoviesTestCase()

	mockService.On("GetMovies", pageStr, sizeStr, mock.Anything).Return(movies, nil)

	resp := utils.PerformRequest(router, "GET", "/movies?page=1&size=5")

	assert.Equal(t, http.StatusOK, resp.Code)
	var response []*models.Movie
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 1)

	expectedMovie := movies[0]
	actualMovie := response[0]
	assert.EqualValues(t, expectedMovie, actualMovie)
}

func TestGetMovieByID(t *testing.T) {
	mockService, router := setupMoviesControllerTest()

	movieID := primitive.NewObjectID().Hex()
	movie := GetMovieByIDTestCase()

	mockService.On("GetMovieByID", movieID, mock.Anything).Return(movie, nil)

	resp := utils.PerformRequest(router, "GET", "/movies/"+movieID)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response models.Movie
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.EqualValues(t, movie, &response)
}
