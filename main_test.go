package main

// This is repositories mock testcases

import (
	"bytes"
	"context"
	"encoding/json"
	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories/mock"
	"example/rest-api-demo/src/routes"
	"example/rest-api-demo/src/services"
	"example/rest-api-demo/src/utils"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockTestSuite struct {
	t                  *testing.T
	moviesRepo         *mock.MockMoviesRepository
	commentsRepo       *mock.MockCommentsRepository
	moviesService      services.MoviesService
	commentsService    services.CommentsService
	moviesController   controllers.MoviesController
	commentsController controllers.CommentsController
	router             *gin.Engine
	ctx                context.Context
}

func setupMockTestSuite(t *testing.T) *MockTestSuite {
	// Initialize mock repositories
	moviesRepo := mock.NewMockMoviesRepository()
	commentsRepo := mock.NewMockCommentsRepository()

	// Initialize services
	moviesService := services.NewMoviesService(moviesRepo)
	commentsService := services.NewCommentsService(commentsRepo)

	// Create configuration
	config := utils.Configuration{
		App: utils.Application{
			Name:    "test_app",
			Timeout: 30,
		},
		Database: utils.DatabaseSetting{
			DbName:      "test_db",
			Collections: []string{"comments", "users", "movies"},
		},
		Server: utils.ServerSettings{
			Port: "8080",
		},
	}

	// Initialize controllers with nil client since we're using mocks
	moviesController := controllers.NewMoviesController(nil, moviesService, config)
	commentsController := controllers.NewCommentsController(nil, commentsService, config)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	controllers := routes.Controllers{
		MoviesController:   moviesController,
		CommentsController: commentsController,
	}
	routes.RegisterRoutes(router, controllers)

	return &MockTestSuite{
		t:                  t,
		moviesRepo:         moviesRepo,
		commentsRepo:       commentsRepo,
		moviesService:      moviesService,
		commentsService:    commentsService,
		moviesController:   moviesController,
		commentsController: commentsController,
		router:             router,
		ctx:                context.Background(),
	}
}

// go test -v -run TestMockMoviesRoutes
func TestMockMoviesRoutes(t *testing.T) {
	suite := setupMockTestSuite(t)

	// go test -v -run TestMockMoviesRoutes/GET_/movies
	t.Run("GET /movies", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/movies", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Movie
		decodeResponse(t, w, &response)
		assert.Len(t, response, 2)
		assert.Equal(t, "Test Movie 1", response[0].Title)
		assert.Equal(t, "Test Movie 2", response[1].Title)
	})

	// go test -v -run TestMockMoviesRoutes/GET_/movies_with_pagination
	t.Run("GET /movies with pagination", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/movies?page=1&size=1", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Movie
		decodeResponse(t, w, &response)
		assert.Len(t, response, 1)
		assert.Equal(t, "Test Movie 1", response[0].Title)
	})

	// go test -v -run "TestMockMoviesRoutes/GET_/movies/:id"
	t.Run("GET /movies/:id", func(t *testing.T) {
		movieID := mock.TestMovieID1.Hex()
		w := makeRequest(t, suite.router, "GET", "/movies/"+movieID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response *models.Movie
		decodeResponse(t, w, &response)
		fmt.Println("response result", response)
		assert.NotNil(t, response)
		assert.Equal(t, "Test Movie 1", response.Title)
	})
}

// go test -v -run TestMockCommentsRoutes
func TestMockCommentsRoutes(t *testing.T) {
	suite := setupMockTestSuite(t)

	// go test -v -run TestMockCommentsRoutes/GET_/comments
	t.Run("GET /comments", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/comments", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Comment
		decodeResponse(t, w, &response)
		assert.Len(t, response, 2)
		assert.Equal(t, "test1@example.com", response[0].Email)
		assert.Equal(t, "test2@example.com", response[1].Email)
	})

	// go test -v -run TestMockCommentsRoutes/GET_/comments/:id
	t.Run("GET /comments/:id", func(t *testing.T) {
		commentID := mock.TestCommentID1.Hex()
		w := makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response *models.Comment
		decodeResponse(t, w, &response)
		assert.NotNil(t, response)
		assert.Equal(t, "test1@example.com", response.Email)
	})

	// go test -v -run TestMockCommentsRoutes/POST_/comments
	t.Run("POST /comments", func(t *testing.T) {
		requestBody := map[string]interface{}{
			"date":     "2025-02-04T23:39:16Z",
			"email":    "damontest@demo.com",
			"movie_id": mock.TestMovieID1.Hex(),
			"name":     "Andrea Le",
			"text":     "This is a test comment for create comment api call",
		}

		w := makeRequest(t, suite.router, "POST", "/comments", requestBody)
		assert.Equal(t, http.StatusCreated, w.Code)

		var response struct {
			Id      string `json:"id"`
			Message string `json:"message"`
		}
		decodeResponse(t, w, &response)
		assert.Equal(t, "Comment created successfully", response.Message)
		assert.NotEmpty(t, response.Id)
	})

	// go test -v -run TestMockCommentsRoutes/PUT_/comments/:id
	t.Run("PUT /comments/:id", func(t *testing.T) {
		commentID := mock.TestCommentID1.Hex()
		requestBody := map[string]interface{}{
			"name":  "Updated Name",
			"email": "updated@example.com",
			"text":  "Updated comment text",
		}

		w := makeRequest(t, suite.router, "PUT", "/comments/"+commentID, requestBody)
		assert.Equal(t, http.StatusOK, w.Code)

		// Verify the update
		w = makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response *models.Comment
		decodeResponse(t, w, &response)
		assert.Equal(t, "Updated Name", response.Name)
		assert.Equal(t, "updated@example.com", response.Email)
		assert.Equal(t, "Updated comment text", response.Text)
	})

	// go test -v -run TestMockCommentsRoutes/DELETE_/comments/:id
	t.Run("DELETE /comments/:id", func(t *testing.T) {
		commentID := mock.TestCommentID1.Hex()
		w := makeRequest(t, suite.router, "DELETE", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		// Verify the deletion
		w = makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

// Helper function to make HTTP requests
func makeRequest(t *testing.T, router *gin.Engine, method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		assert.NoError(t, err)
	}

	req, err := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	assert.NoError(t, err)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// Helper function to decode response
func decodeResponse(t *testing.T, w *httptest.ResponseRecorder, v interface{}) {
	err := json.NewDecoder(w.Body).Decode(v)
	assert.NoError(t, err)
}
