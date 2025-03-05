package controllers

import (
	"example/rest-api-demo/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetMoviesTestCase() []*models.Movie {
	return []*models.Movie{
		{
			ID:               primitive.NewObjectID(),
			Awards:           models.Awards{Nominations: 3, Text: "3 nominations", Wins: 1},
			Cast:             []string{"Actor 1", "Actor 2"},
			Countries:        []string{"USA"},
			Directors:        []string{"Director 1"},
			FullPlot:         "Full plot of the movie",
			Genres:           []string{"Action", "Adventure"},
			IMDb:             models.IMDb{ID: 123456, Rating: 7.8, Votes: 1000},
			NumMflixComments: 10,
			Plot:             "Short plot of the movie",
			Rated:            "PG-13",
			Runtime:          120,
			Title:            "Movie Title",
			Type:             "movie",
			Year:             2021,
		},
	}
}

func GetMovieByIDTestCase() *models.Movie {
	return &models.Movie{
		ID:               primitive.NewObjectID(),
		Awards:           models.Awards{Nominations: 3, Text: "3 nominations", Wins: 1},
		Cast:             []string{"Actor 1", "Actor 2"},
		Countries:        []string{"USA"},
		Directors:        []string{"Director 1"},
		FullPlot:         "Full plot of the movie",
		Genres:           []string{"Action", "Adventure"},
		IMDb:             models.IMDb{ID: 123456, Rating: 7.8, Votes: 1000},
		NumMflixComments: 10,
		Plot:             "Short plot of the movie",
		Rated:            "PG-13",
		Runtime:          120,
		Title:            "Movie Title",
		Type:             "movie",
		Year:             2021,
	}
}
