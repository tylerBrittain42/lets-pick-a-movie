package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/tylerBrittain42/lets-pick-a-movie/internal/pageTemplate"
)

func (app *application) handleSingleMovie(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("movieName")
	unescaped, err := url.PathUnescape(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error handling escape characters in url: %s", err)
		return
	}

	movie, err := app.queries.GetMovie(r.Context(), unescaped)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error grabbing movie: %s", err)
		return
	}
	pageTemplate.PageMovie(movie).Render(r.Context(), w)
}

// todo: paginate movies

func (app *application) handleAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.queries.GetMovies(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to query movies: %s", err)
		return

	}

	_ = movies
	_ = pageTemplate.PageMovies(movies[:10]).Render(r.Context(), w)
	// a_ := movies[0]

}
