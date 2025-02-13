package repositories

import (
	"context"
	"log"

	"example/rest-api-demo/src/models"
	"example/rest-api-demo/src/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MoviesRepository interface {
	GetMovies(page, size int, ctx context.Context) ([]*models.Movie, error)
	GetMovieByID(objID primitive.ObjectID, ctx context.Context) (*models.Movie, error)
}

type moviesRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewMovieMongodbRepo(config *utils.Configuration, client *mongo.Client) MoviesRepository {
	return &moviesRepository{config: config, client: client}
}

func (mcr moviesRepository) GetMovies(page, size int, ctx context.Context) ([]*models.Movie, error) {
	findOptions := options.Find()
	findOptions.SetLimit(int64(size))
	findOptions.SetSkip(int64((page - 1) * size))

	collection := mcr.client.Database(mcr.config.Database.DbName).Collection(mcr.config.Database.Collections[2])
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var movies []*models.Movie

	for cursor.Next(ctx) {
		var movie models.Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Println("Error decoding movie:", err)
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (mcr moviesRepository) GetMovieByID(objID primitive.ObjectID, ctx context.Context) (*models.Movie, error) {
	var movie *models.Movie
	collection := mcr.client.Database(mcr.config.Database.DbName).Collection(mcr.config.Database.Collections[2])

	err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Movie not found")
		} else {
			log.Println("Error finding document:", err)
		}
		return movie, err
	}

	return movie, nil
}
