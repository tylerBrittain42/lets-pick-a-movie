package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleSingleMovie(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("movieName")
	w.Write([]byte(fmt.Sprintf("stub for %s\n", name)))
}

// todo: paginate movies

func (app *application) handleAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.queries.GetMovies(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to query movies: %s", err)
		return

	}
	w.Write([]byte(fmt.Sprintf("%v", movies)))
}
