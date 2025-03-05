package mongodb_repo

import (
	"context"
	"log"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/repositories"
	"example/rest-api-demo/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentsRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewCommentMongodbRepo(config *utils.Configuration, client *mongo.Client) repositories.CommentsRepository {
	return &commentsRepository{config: config, client: client}
}

func (cr commentsRepository) GetComments(findOptions *options.FindOptions, ctx context.Context) ([]*models.Comment, error) {
	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var comments []*models.Comment

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			log.Println("Error decoding comment:", err)
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (cr commentsRepository) GetCommentByID(idStr string, ctx context.Context) (*models.Comment, error) {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	var comment *models.Comment
	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Comment not found")
		} else {
			log.Println("Error finding document:", err)
		}
		return comment, err
	}

	return comment, nil
}

func (cr commentsRepository) DeleteCommentByID(idStr string, ctx context.Context) error {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}

	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (cr commentsRepository) UpdateCommentByID(idStr string, updateData bson.M, ctx context.Context) error {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}

	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])

	update := bson.M{"$set": updateData}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (cr commentsRepository) CreateComment(comment *models.Comment, ctx context.Context) (string, error) {
	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])

	result, err := collection.InsertOne(ctx, comment)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (cr commentsRepository) GetCommentsForMovie(findOptions *options.FindOptions, idStr string, ctx context.Context) ([]*models.Comment, error) {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	collection := cr.client.Database(cr.config.Database.DbName).Collection(cr.config.Database.Collections[0])
	cursor, err := collection.Find(ctx, bson.M{"movie_id": objID}, findOptions)
	if err != nil {
		return nil, err
	}

	var comments []*models.Comment

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			log.Println("Error decoding comment:", err)
			return nil, err
		}

		log.Println("Comment:", &comment)
		comments = append(comments, &comment)
	}

	return comments, nil
}
