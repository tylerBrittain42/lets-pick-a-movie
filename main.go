package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

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
		Handler:      middlewareLogReq(mux),
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

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func middlewareLogReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{w, http.StatusOK}
		next.ServeHTTP(&rec, r)
		reqInfo := (fmt.Sprintf("%s %s status:%d", r.Method, r.URL.RequestURI(), rec.status))
		fmt.Printf("Handled: %s\n", reqInfo)
	})
}
