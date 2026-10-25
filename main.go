package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err.Error())
	}

	port := os.Getenv("PORT")

	users := http.NewServeMux()
	users.HandleFunc("GET /login", handleLogin)
	users.HandleFunc("GET /register", handleRegisterUser)
	users.HandleFunc("GET /{userName}", handleUserProfile)
	users.HandleFunc("GET /compare", handleOverlap)

	movies := http.NewServeMux()
	movies.HandleFunc("GET /{movieName}", handleSingleMovie)
	movies.HandleFunc("GET /{$}", handleAllMovies)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /uploadCsv", handleUploadCsv)
	mux.HandleFunc("GET /{$}", handleHome)

	mux.Handle("/users/", http.StripPrefix("/users", users))
	mux.Handle("/movies/", http.StripPrefix("/movies", movies))

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")

}

func handleUploadCsv(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("upload stub"))
}

// will user query parameters
func handleOverlap(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("overlap stub"))
}
