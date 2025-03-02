package services

import (
	"context"
	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories/mongodb_repo"
	"example/rest-api-demo/src/utils"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentsService interface {
	GetComments(pageStr, sizeStr string, ctx context.Context) ([]*models.Comment, error)
	GetCommentByID(id string, ctx context.Context) (*models.Comment, error)
	DeleteCommentByID(id string, ctx context.Context) error
	UpdateCommentByID(id string, updateData bson.M, ctx context.Context) error
	CreateComment(comment models.Comment, ctx context.Context) (string, error)
}

type commentsService struct {
	cr mongodb_repo.CommentsRepository
}

func NewCommentsService(cr mongodb_repo.CommentsRepository) CommentsService {
	return &commentsService{cr: cr}
}

func (cs *commentsService) GetComments(pageStr, sizeStr string, ctx context.Context) ([]*models.Comment, error) {
	pagination := utils.GetPaginationParams(pageStr, sizeStr)

	findOptions := options.Find()
	findOptions.SetLimit(int64(pagination.Size))
	findOptions.SetSkip(int64((pagination.Page - 1) * pagination.Size))

	return cs.cr.GetComments(findOptions, ctx)
}

func (cs *commentsService) GetCommentByID(idStr string, ctx context.Context) (*models.Comment, error) {
	log.Println("Getting comment by ID", idStr)
	return cs.cr.GetCommentByID(idStr, ctx)
}

func (cs *commentsService) DeleteCommentByID(idStr string, ctx context.Context) error {
	log.Println("Deleting comment by ID", idStr)
	return cs.cr.DeleteCommentByID(idStr, ctx)
}

func (cs *commentsService) UpdateCommentByID(idStr string, updateData bson.M, ctx context.Context) error {
	log.Println("Updating comment by ID", idStr)
	return cs.cr.UpdateCommentByID(idStr, updateData, ctx)
}

func (cs *commentsService) CreateComment(comment models.Comment, ctx context.Context) (string, error) {
	log.Println("Creating comment")
	return cs.cr.CreateComment(comment, ctx)
}
