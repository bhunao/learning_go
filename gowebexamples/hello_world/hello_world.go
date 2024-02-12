package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)
func handler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())

	minParam := r.URL.Query().Get("min")
	maxParam := r.URL.Query().Get("max")

	min, _ := strconv.Atoi(minParam)
	max, _ := strconv.Atoi(maxParam)

	randomNumber := rand.Intn(max-min+1) + min
	fmt.Fprintf(w, "Random Number: %d", randomNumber)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
