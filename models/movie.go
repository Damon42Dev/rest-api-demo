package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Awards struct {
	Nominations int    `bson:"nominations" json:"nominations"`
	Text        string `bson:"text" json:"text"`
	Wins        int    `bson:"wins" json:"wins"`
}

type IMDb struct {
	ID     int     `bson:"id" json:"id"`
	Rating float64 `bson:"rating" json:"rating"`
	Votes  int     `bson:"votes" json:"votes"`
}

type Viewer struct {
	Meter      int     `bson:"meter" json:"meter"`
	NumReviews int     `bson:"numReviews" json:"numReviews"`
	Rating     float64 `bson:"rating" json:"rating"`
}

type Tomatoes struct {
	LastUpdated time.Time `bson:"lastUpdated" json:"lastUpdated"`
	Viewer      Viewer    `bson:"viewer" json:"viewer"`
}

type Movie struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Awards    Awards             `bson:"awards" json:"awards"`
	Cast      []string           `bson:"cast" json:"cast"`
	Countries []string           `bson:"countries" json:"countries"`
	Directors []string           `bson:"directors" json:"directors"`
	FullPlot  string             `bson:"fullplot" json:"fullplot"`
	Genres    []string           `bson:"genres" json:"genres"`
	IMDb      IMDb               `bson:"imdb" json:"imdb"`
	// LastUpdated      CustomTime         `bson:"lastupdated" json:"lastupdated"`
	NumMflixComments int       `bson:"num_mflix_comments" json:"num_mflix_comments"`
	Plot             string    `bson:"plot" json:"plot"`
	Rated            string    `bson:"rated" json:"rated"`
	Released         time.Time `bson:"released" json:"released"`
	Runtime          int       `bson:"runtime" json:"runtime"`
	Title            string    `bson:"title" json:"title"`
	Tomatoes         Tomatoes  `bson:"tomatoes" json:"tomatoes"`
	Type             string    `bson:"type" json:"type"`
	Year             int       `bson:"year" json:"year"`
}
