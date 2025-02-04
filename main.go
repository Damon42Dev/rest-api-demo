package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func main() {
	db = Connect()

	server := gin.Default()
	server.GET("/movies", GetMovies)
	server.GET("/comments", GetComments)
	server.GET("/comments/:id", GetCommentByID)
	server.POST("/comments", CreateComment)
	server.PUT("/comments/:id", UpdateCommentByID)
	server.DELETE("/comments/:id", DeleteCommentByID)

	server.GET("/movies/:id", GetMovieByID)
	server.Run(":8080")
}

func Connect() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client.Database("sample_mflix")
}

func GetMovies(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.Atoi(pageStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil || page < 1 {
		page = 1
	}

	if err != nil || size < 1 {
		size = 10
	}

	limit := int64(10)
	skip := int64((page - 1) * 10)

	collection := db.Collection("movies")
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding documents: %s", err)})
		return
	}
	defer cursor.Close(context.Background())

	var movies []bson.M
	if err := cursor.All(context.Background(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding documents: %s", err)})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func GetComments(c *gin.Context) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	page, err := strconv.Atoi(pageStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil || page < 1 {
		page = 1
	}

	if err != nil || size < 1 {
		size = 1
	}

	limit := int64(size)
	skip := int64((page - 1) * 10)

	collection := db.Collection("comments")
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding documents: %s", err)})
		return
	}
	defer cursor.Close(context.Background())

	var comments []bson.M
	if err := cursor.All(context.Background(), &comments); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error decoding documents: %s", err)})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func GetCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := db.Collection("comments")
	var comment bson.M
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&comment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding document: %s", err)})
		}
		return
	}

	c.JSON(http.StatusOK, comment)
}

func GetMovieByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := db.Collection("movies")
	var movie bson.M
	err = collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error finding document: %s", err)})
		}
		return
	}

	c.JSON(http.StatusOK, movie)
}

func DeleteCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := db.Collection("comments")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error deleting document: %s", err)})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func UpdateCommentByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateData bson.M
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	collection := db.Collection("comments")
	update := bson.M{"$set": updateData}
	result, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating document: %s", err)})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func CreateComment(c *gin.Context) {
	var comment bson.M
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	collection := db.Collection("comments")
	result, err := collection.InsertOne(context.Background(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error inserting document: %s", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "id": result.InsertedID})
}
