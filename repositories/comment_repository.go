package repositories

import (
	"context"
	"example/rest-api-demo/config"
	"example/rest-api-demo/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetComments(pageStr, sizeStr string) ([]models.Comment, error) {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)

	if err != nil || size < 1 {
		size = 10
	}

	limit := int64(10)
	skip := int64((page - 1) * 10)

	var comments []models.Comment
	collection := config.GetCollection("comments")
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Println("Error fetching comments:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var comment models.Comment
		if err := cursor.Decode(&comment); err != nil {
			log.Println("Error decoding comment:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func GetCommentByID(objID primitive.ObjectID) (models.Comment, error) {
	var comment models.Comment
	collection := config.GetCollection("comments")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&comment)
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

func DeleteCommentByID(objID primitive.ObjectID) error {
	collection := config.GetCollection("comments")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func UpdateCommentByID(objID primitive.ObjectID, updateData bson.M) error {
	collection := config.GetCollection("comments")
	update := bson.M{"$set": updateData}
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
