package main

// This is local installed mongoDB testcases

import (
	"bytes"
	"context"
	"encoding/json"
	"example/rest-api-demo/src/controllers"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/repositories/mongodb_repo"
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

type TestSuite struct {
	t                  *testing.T
	testDB             *utils.TestDBClient
	moviesRepo         repositories.MoviesRepository
	commentsRepo       repositories.CommentsRepository
	moviesService      services.MoviesService
	commentsService    services.CommentsService
	moviesController   controllers.MoviesController
	commentsController controllers.CommentsController
	router             *gin.Engine
	ctx                context.Context
}

func setupTestSuite(t *testing.T) *TestSuite {
	// Setup test database
	testDB := utils.SetupTestDB(t)
	t.Cleanup(func() {
		utils.TeardownTestDB(t, testDB.Client)
	})

	// Initialize repositories
	moviesRepo := mongodb_repo.NewMovieMongodbRepo(testDB.Config, testDB.Client)
	commentsRepo := mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

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

	// Initialize controllers
	moviesController := controllers.NewMoviesController(testDB.Client, moviesService, config)
	commentsController := controllers.NewCommentsController(testDB.Client, commentsService, config)

	// Setup router
	gin.SetMode(gin.TestMode)
	router := gin.New()
	controllers := routes.Controllers{
		MoviesController:   moviesController,
		CommentsController: commentsController,
	}
	routes.RegisterRoutes(router, controllers)

	return &TestSuite{
		t:                  t,
		testDB:             testDB,
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

// HTTP Layer Tests
// go test -v -run TestMoviesRoutes
func TestMoviesRoutes(t *testing.T) {
	suite := setupTestSuite(t)

	// go test -v -run TestMoviesRoutes/GET_/movies
	t.Run("GET /movies", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/movies", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Movie
		decodeResponse(t, w, &response)
		assert.Len(t, response, 2)
		fmt.Printf("Response: %+v\n", response[0])
		assert.Equal(t, "Test Movie 1", response[0].Title)
		assert.Equal(t, "Test Movie 2", response[1].Title)
	})

	// go test -v -run "TestMoviesRoutes/GET_/movies_with_pagination"
	t.Run("GET /movies with pagination", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/movies?page=1&size=1", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Movie
		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response[0])
		assert.Len(t, response, 1)
		assert.Equal(t, "Test Movie 1", response[0].Title)
	})

	// go test -v -run "TestMoviesRoutes/GET_/movies/:id"
	t.Run("GET /movies/:id", func(t *testing.T) {
		movieID := "60c72b2f9b1e8a5d6c8b4567"
		w := makeRequest(t, suite.router, "GET", "/movies/"+movieID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response *models.Movie
		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response)
		assert.NotNil(t, response)
		assert.Equal(t, "Test Movie 1", response.Title)
	})
}

// go test -v -run TestCommentsRoutes
func TestCommentsRoutes(t *testing.T) {
	suite := setupTestSuite(t)

	// go test -v -run "TestCommentsRoutes/GET_/comments"
	t.Run("GET /comments", func(t *testing.T) {
		w := makeRequest(t, suite.router, "GET", "/comments", nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response []*models.Comment
		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response[0])
		assert.Len(t, response, 2)
	})

	// go test -v -run "TestCommentsRoutes/GET_/comments/:id"
	t.Run("GET /comments/:id", func(t *testing.T) {
		commentID := "60c72b2f9b1e8a5d6c8b4567"
		w := makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response *models.Comment
		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response)
		assert.NotNil(t, response)
		assert.Equal(t, "test1@example.com", response.Email)
	})

	// go test -v -run "TestCommentsRoutes/POST_/comments"
	t.Run("POST /comments", func(t *testing.T) {
		// Create request body with correct field name
		requestBody := map[string]interface{}{
			"date":     "2025-02-04T23:39:16Z",
			"email":    "damontest@demo.com",
			"movie_id": "573a1390f29313caabcd418c",
			"name":     "Andrea Le updated",
			"text":     "This is a test comment for create comment api call updated",
		}

		// Log the request body
		body, _ := json.Marshal(requestBody)
		fmt.Printf("Request body: %s\n", string(body))

		w := makeRequest(t, suite.router, "POST", "/comments", requestBody)
		assert.Equal(t, http.StatusCreated, w.Code)

		var response struct {
			Id      string `json:"id"`
			Message string `json:"message"`
		}

		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response)
		assert.Equal(t, "Comment created successfully", response.Message)
		assert.NotEmpty(t, response.Id)
	})

	// go test -v -run "TestCommentsRoutes/PUT_/comments/:id"
	t.Run("PUT /comments/:id", func(t *testing.T) {
		commentID := "60c72b2f9b1e8a5d6c8b4567"

		// Create update request body
		requestBody := map[string]interface{}{
			"text": "This is an updated comment text",
			"name": "Updated Name",
		}

		// Log the request body
		body, _ := json.Marshal(requestBody)
		fmt.Printf("Request body: %s\n", string(body))

		w := makeRequest(t, suite.router, "PUT", "/comments/"+commentID, requestBody)
		assert.Equal(t, http.StatusOK, w.Code)

		var response struct {
			Message string `json:"message"`
		}

		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response)
		assert.Equal(t, "Comment updated successfully", response.Message)

		// Verify the update by fetching the comment
		w = makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var comment *models.Comment
		decodeResponse(t, w, &comment)
		fmt.Printf("Updated comment: %+v\n", comment)
		assert.Equal(t, "This is an updated comment text", comment.Text)
		assert.Equal(t, "Updated Name", comment.Name)
	})

	// go test -v -run "TestCommentsRoutes/DELETE_/comments/:id"
	t.Run("DELETE /comments/:id", func(t *testing.T) {
		commentID := "60c72b2f9b1e8a5d6c8b4567"

		// First verify the comment exists
		w := makeRequest(t, suite.router, "GET", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		// Delete the comment
		w = makeRequest(t, suite.router, "DELETE", "/comments/"+commentID, nil)
		assert.Equal(t, http.StatusOK, w.Code)

		var response struct {
			Message string `json:"message"`
		}

		decodeResponse(t, w, &response)
		fmt.Printf("Response: %+v\n", response)
		assert.Equal(t, "Comment deleted successfully", response.Message)

		// Verify the comment is deleted by trying to fetch it
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
