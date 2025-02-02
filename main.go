package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"

	"context"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	db := Connect()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	// http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
	// 	t := Time{CurrentTime: time.Now().String()}
	// 	json.NewEncoder(w).Encode(t)
	// })

	http.HandleFunc("/collection/users", func(w http.ResponseWriter, r *http.Request) {
		collection := db.Collection("users")

		cursor, err := collection.Find(context.Background(), bson.D{{}})
		if err != nil {
			http.Error(w, fmt.Sprintf("Error finding documents: %s", err), http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.Background())

		var documents []bson.M
		if err := cursor.All(context.Background(), &documents); err != nil {
			http.Error(w, fmt.Sprintf("Error decoding documents: %s", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(documents)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func Connect() *mongo.Database {
	// Find .evn
	err := godotenv.Load("/opt/.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Get value from .env
	MONGODB_URI := os.Getenv("MONGODB_URI")

	// Connect to the database.
	clientOption := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
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
