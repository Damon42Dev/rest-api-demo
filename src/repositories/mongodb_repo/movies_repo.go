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

type moviesRepository struct {
	client *mongo.Client
	config *utils.Configuration
}

func NewMovieMongodbRepo(config *utils.Configuration, client *mongo.Client) repositories.MoviesRepository {
	return &moviesRepository{config: config, client: client}
}

func (mr moviesRepository) GetMovies(findOptions *options.FindOptions, ctx context.Context) ([]*models.Movie, error) {
	collection := mr.client.Database(mr.config.Database.DbName).Collection(mr.config.Database.Collections[2])
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

func (mr moviesRepository) GetMovieByID(idStr string, ctx context.Context) (*models.Movie, error) {
	objID, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return nil, err
	}

	var movie *models.Movie
	collection := mr.client.Database(mr.config.Database.DbName).Collection(mr.config.Database.Collections[2])

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&movie)
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
