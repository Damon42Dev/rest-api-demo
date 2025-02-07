package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Date    string             `bson:"date" json:"date"`
	Email   string             `bson:"email" json:"email"`
	MovieID primitive.ObjectID `bson:"movie_id" json:"movie_id"`
	Name    string             `bson:"name" json:"name"`
	Text    string             `bson:"text" json:"text"`
}
