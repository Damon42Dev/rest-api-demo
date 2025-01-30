package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

type Time struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := Time{CurrentTime: time.Now().String()}
		json.NewEncoder(w).Encode(t)
	})

	fmt.Println("Server is running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
