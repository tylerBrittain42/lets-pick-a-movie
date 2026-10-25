package main

import (
	"fmt"
	"net/http"
)

func handleSingleMovie(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("movieName")
	w.Write([]byte(fmt.Sprintf("stub for %s\n", name)))
}
func handleAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("stub for all movies"))
}
