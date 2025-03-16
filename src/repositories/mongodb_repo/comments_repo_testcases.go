package mongodb_repo

import (
	"example/rest-api-demo/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCommentsTestCase() []*models.Comment {
	return []*models.Comment{
		{
			ID:      primitive.NewObjectID(),
			Date:    primitive.NewDateTimeFromTime(time.Now()),
			Email:   "test1@example.com",
			MovieID: primitive.NewObjectID(),
			Name:    "Test User 1",
			Text:    "This is a test comment 1",
		},
		{
			ID:      primitive.NewObjectID(),
			Date:    primitive.NewDateTimeFromTime(time.Now()),
			Email:   "test2@example.com",
			MovieID: primitive.NewObjectID(),
			Name:    "Test User 2",
			Text:    "This is a test comment 2",
		},
	}
}

func GetCommentByIDTestCase() *models.Comment {
	return &models.Comment{
		ID:      primitive.NewObjectID(),
		Date:    primitive.NewDateTimeFromTime(time.Now()),
		Email:   "test@example.com",
		MovieID: primitive.NewObjectID(),
		Name:    "Test User",
		Text:    "This is a test comment",
	}
}
