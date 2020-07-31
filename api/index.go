package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var bros = []string{
	"broast beef.",
	"Broe Biden.",
}

var greetings = []string{
	"Hello",
	"Hi",
	"Good day",
	"Greetings",
}

// Response is an API response.
type Response struct {
	Message string `json:"message"`
}

// Handler handles a serverless invocation.
func Handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())

	bro := bros[rand.Intn(len(bros))]
	greeting := greetings[rand.Intn(len(greetings))]
	message := fmt.Sprintf("%s, %s.", greeting, bro)

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	response := Response{Message: message}
	json.NewEncoder(w).Encode(response)
}
