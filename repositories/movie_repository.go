package repositories

import (
	"context"
	"example/rest-api-demo/config"
	"example/rest-api-demo/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMovies(pageStr, sizeStr string) ([]models.Movie, error) {
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

	var movies []models.Movie
	collection := config.GetCollection("movies")
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(skip)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Println("Error fetching movies:", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Error decoding movie:", err)
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
