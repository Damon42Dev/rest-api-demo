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
	"go.mongodb.org/mongo-driver/bson"
)

// MockCommentsService is a mock implementation of the CommentsService interface
type MockCommentsService struct {
	mock.Mock
}

func (m *MockCommentsService) GetComments(pageStr, sizeStr string, ctx context.Context) ([]*models.Comment, error) {
	args := m.Called(pageStr, sizeStr, ctx)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func (m *MockCommentsService) GetCommentByID(id string, ctx context.Context) (*models.Comment, error) {
	args := m.Called(id, ctx)
	return args.Get(0).(*models.Comment), args.Error(1)
}

func (m *MockCommentsService) DeleteCommentByID(id string, ctx context.Context) error {
	args := m.Called(id, ctx)
	return args.Error(0)
}

func (m *MockCommentsService) UpdateCommentByID(id string, updateData bson.M, ctx context.Context) error {
	args := m.Called(id, updateData, ctx)
	return args.Error(0)
}

func (m *MockCommentsService) CreateComment(comment models.Comment, ctx context.Context) (string, error) {
	args := m.Called(comment, ctx)
	return args.String(0), args.Error(1)
}

func (m *MockCommentsService) GetCommentsForMovie(pageStr, sizeStr, idStr string, ctx context.Context) ([]*models.Comment, error) {
	args := m.Called(pageStr, sizeStr, idStr, ctx)
	return args.Get(0).([]*models.Comment), args.Error(1)
}

func setupCommentsControllerTest() (*MockCommentsService, *gin.Engine) {
	mockService := new(MockCommentsService)
	config := utils.Configuration{
		App: utils.Application{
			Timeout: 5,
		},
	}
	controller := NewCommentsController(nil, mockService, config)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/comments", controller.GetComments)
	router.GET("/comments/:id", controller.GetCommentByID)
	router.DELETE("/comments/:id", controller.DeleteCommentByID)
	router.PUT("/comments/:id", controller.UpdateCommentByID)
	router.POST("/comments", controller.CreateComment)
	router.GET("/movies/:id/comments", controller.GetCommentsForMovie)

	return mockService, router
}

func TestGetComments(t *testing.T) {
	mockService, router := setupCommentsControllerTest()

	pageStr := "1"
	sizeStr := "5"
	comments := GetCommentsTestCase()

	mockService.On("GetComments", pageStr, sizeStr, mock.Anything).Return(comments, nil)

	resp := utils.PerformRequest(router, "GET", "/comments?page=1&size=5", nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response []*models.Comment
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)

	expectedComment := comments[0]
	actualComment := response[0]
	assert.EqualValues(t, expectedComment, actualComment)
}

func TestGetCommentByID(t *testing.T) {
	mockService, router := setupCommentsControllerTest()

	commentID := "60c72b2f9b1e8a5d6c8b4567" // Specific string value for testing
	comment := GetCommentByIDTestCase()

	mockService.On("GetCommentByID", commentID, mock.Anything).Return(comment, nil)

	resp := utils.PerformRequest(router, "GET", "/comments/"+commentID, nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response models.Comment
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.EqualValues(t, comment, &response)
}

func TestDeleteCommentByID(t *testing.T) {
	mockService, router := setupCommentsControllerTest()

	commentID := "60c72b2f9b1e8a5d6c8b4567" // Specific string value for testing

	mockService.On("DeleteCommentByID", commentID, mock.Anything).Return(nil)

	resp := utils.PerformRequest(router, "DELETE", "/comments/"+commentID, nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Comment deleted successfully", response["message"])
}

func TestUpdateCommentByID(t *testing.T) {
	mockService, router := setupCommentsControllerTest()

	commentID := "60c72b2f9b1e8a5d6c8b4567" // Specific string value for testing
	updateData := bson.M{"text": "Updated Comment"}

	mockService.On("UpdateCommentByID", commentID, updateData, mock.Anything).Return(nil)

	resp := utils.PerformRequest(router, "PUT", "/comments/"+commentID, updateData)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Comment updated successfully", response["message"])
}

// func TestCreateComment(t *testing.T) {
// 	mockService, router := setupCommentsControllerTest()

// 	comment := GetCommentByIDTestCase()
// 	commentID := primitive.NewObjectID()

// 	mockService.On("CreateComment", comment, mock.Anything).Return(commentID.Hex(), nil)

// 	resp := utils.PerformRequest(router, "POST", "/comments", comment)

// 	assert.Equal(t, http.StatusCreated, resp.Code)
// 	var response map[string]interface{}
// 	err := json.Unmarshal(resp.Body.Bytes(), &response)
// 	assert.NoError(t, err)
// 	assert.Equal(t, "Comment created successfully", response["message"])
// 	idStr, ok := response["id"].(string)
// 	assert.True(t, ok)
// 	assert.Equal(t, commentID.Hex(), idStr)
// }

func TestGetCommentsForMovie(t *testing.T) {
	mockService, router := setupCommentsControllerTest()

	movieID := "60c72b2f9b1e8a5d6c8b4567" // Specific string value for testing
	pageStr := "1"
	sizeStr := "5"
	comments := GetCommentsTestCase()

	mockService.On("GetCommentsForMovie", pageStr, sizeStr, movieID, mock.Anything).Return(comments, nil)

	resp := utils.PerformRequest(router, "GET", "/movies/"+movieID+"/comments?page=1&size=5", nil)

	assert.Equal(t, http.StatusOK, resp.Code)
	var response []*models.Comment
	err := json.Unmarshal(resp.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Len(t, response, 2)

	expectedComment := comments[0]
	actualComment := response[0]
	assert.EqualValues(t, expectedComment, actualComment)
}
