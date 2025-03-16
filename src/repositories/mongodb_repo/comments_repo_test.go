package mongodb_repo_test

import (
	"context"
	"testing"
	"time"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/repositories/mongodb_repo"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var commentsRepo repositories.CommentsRepository

func TestGetComments(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test GetComments
	findOptions := options.Find()
	result, err := commentsRepo.GetComments(findOptions, context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "test1@example.com", result[0].Email)
	assert.Equal(t, "test2@example.com", result[1].Email)
}

func TestGetCommentByID(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test GetCommentByID
	result, err := commentsRepo.GetCommentByID("60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test1@example.com", result.Email)
}

func TestDeleteCommentByID(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test DeleteCommentByID
	err := commentsRepo.DeleteCommentByID("60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.NoError(t, err)

	// Verify deletion
	result, err := commentsRepo.GetCommentByID("60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateCommentByID(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test UpdateCommentByID
	updateData := bson.M{"text": "Updated Comment"}
	err := commentsRepo.UpdateCommentByID("60c72b2f9b1e8a5d6c8b4567", updateData, context.Background())
	assert.NoError(t, err)

	// Verify update
	result, err := commentsRepo.GetCommentByID("60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Comment", result.Text)
}

func TestCreateComment(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test CreateComment
	comment := &models.Comment{
		ID:      primitive.NewObjectID(),
		Date:    primitive.NewDateTimeFromTime(time.Now()),
		Email:   "test3@example.com",
		MovieID: primitive.NewObjectID(),
		Name:    "Test User 3",
		Text:    "This is a test comment 3",
	}
	id, err := commentsRepo.CreateComment(comment, context.Background())
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	// Verify creation
	result, err := commentsRepo.GetCommentByID(id, context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "test3@example.com", result.Email)
}

func TestGetCommentsForMovie(t *testing.T) {
	// Setup
	testDB := mongodb_repo.SetupTestDB(t)
	defer mongodb_repo.TeardownTestDB(t, testDB.Client)
	commentsRepo = mongodb_repo.NewCommentMongodbRepo(testDB.Config, testDB.Client)

	// Test GetCommentsForMovie
	findOptions := options.Find()
	result, err := commentsRepo.GetCommentsForMovie(findOptions, "60c72b2f9b1e8a5d6c8b4567", context.Background())
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "test1@example.com", result[0].Email)
	assert.Equal(t, "test2@example.com", result[1].Email)
}
