package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var bros = []string{
	"broast beef",
	"Broe Biden",
	"Bro Montana",
	"Brosiah",
	"Brozo the Clown",
	"Angelina Brolie",
	"Marco Brolo",
	"Plácido Bromingo",
	"Vincent Van Brogh",
	"Sandy Broufax",
	"Lebrona Helmsley",
	"Brobi-Wan Kenobi",
	"Brometheus",
	"Elvis Costellbro",
	"Stephen Brolbert",
	"Nabroleon Bronaparte",
	"Evander Brolyfield",
	"Mario Brotali",
	"Terrell Browens",
	"Tony Bromo",
	"Pete Brose",
	"Brony Soprano",
	"Bro Jackson",
	"Bronus Wagner",
	"Emilibro Estevez",
	"Pablo Picassbro",
	"Broto Baggins",
	"Willem Dabroe",
	"Francis Ford Broppola",
	"John Broltrane",
	"Teddy Broosevelt",
	"Brogi Berra",
	"Tim Tebro",
	"Edgar Allen Bro",
	"Rice o Broni",
	"Pete Brozelle",
	"C-3PBro",
	"Bromer Simpson",
	"Nicolas Sarbrozy",
	"Broprah Winfrey",
	"Axl Brose",
	"Sherlock Brolmes",
	"Bro Derek",
	"Bro Diddley",
	"Yo-Yo Brah",
	"Brosé Feliciano",
	"Jon Favbreau",
	"Mephistophbroles",
	"Diego Marabrona",
	"G.I. Bro",
	"Fats Bromino",
	"Brollie Fingers",
	"Ringbro Starr",
	"Kurt Brobain",
	"Alec Broldwin",
	"Dom DiMaggibro",
	"Shaquille brO’Neal",
	"Muggsy Brogues",
}

var greetings = []string{
	"Hello",
	"Hi",
	"Good day",
	"Greetings",
	"So long",
	"Goodbye",
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
