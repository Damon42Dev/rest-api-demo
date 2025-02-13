package repositories

import (
	"context"
	"log"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentsRepository interface {
	// Add(appDoc entity.AppDoc, ctx context.Context) (primitive.ObjectID, error)
	GetComments(page, size int, ctx context.Context) ([]*models.Comment, error)
	// GetById(oId primitive.ObjectID, ctx context.Context) (*entity.AppDoc, error)
	// Delete(oId primitive.ObjectID, ctx context.Context) (int64, error)
}

type commentsRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewCommentMongodbRepo(config *utils.Configuration, client *mongo.Client) CommentsRepository {
	return &commentsRepository{config: config, client: client}
}

func (mcr commentsRepository) GetComments(page, size int, ctx context.Context) ([]*models.Comment, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(size))
	findOptions.SetSkip(int64((page - 1) * size))

	collection := mcr.client.Database(mcr.config.Database.DbName).Collection(mcr.config.Database.Collections[0])
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

// func (mcr MongoDBCommentRepo) GetCommentByID(objID primitive.ObjectID) (models.Comment, error) {
// 	var comment models.Comment
// 	collection := config.GetCollection("comments")

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&comment)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			log.Println("Comment not found")
// 		} else {
// 			log.Println("Error finding document:", err)
// 		}
// 		return comment, err
// 	}

// 	return comment, nil
// }

// func DeleteCommentByID(objID primitive.ObjectID) error {
// 	collection := config.GetCollection("comments")
// 	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
// 	if err != nil {
// 		return err
// 	}

// 	if result.DeletedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// func UpdateCommentByID(objID primitive.ObjectID, updateData bson.M) error {
// 	collection := config.GetCollection("comments")
// 	update := bson.M{"$set": updateData}
// 	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
// 	if err != nil {
// 		return err
// 	}

// 	if result.MatchedCount == 0 {
// 		return mongo.ErrNoDocuments
// 	}

// 	return nil
// }

// func CreateComment(comment models.Comment) (primitive.ObjectID, error) {
// 	collection := config.GetCollection("comments")
// 	result, err := collection.InsertOne(context.Background(), comment)
// 	if err != nil {
// 		return primitive.NilObjectID, err
// 	}

// 	return result.InsertedID.(primitive.ObjectID), nil
// }
