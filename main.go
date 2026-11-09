package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "modernc.org/sqlite"

	"github.com/joho/godotenv"
	"github.com/tylerBrittain42/lets-pick-a-movie/internal/database"
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

	//todo: add something to handle unset env
	// todo: add local mem fallback
	// todo: add these to plan
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DATABASE_URL")
	platform := os.Getenv("PLATFORM")

	fmt.Print(dbUrl)
	db, err := sql.Open("sqlite", dbUrl)
	if err != nil {
		log.Fatalf("Cannot access sqlite %v\n", err)
	}
	dbQueries := database.New(db)
	app := application{queries: dbQueries, platform: platform}

	users := http.NewServeMux()
	users.HandleFunc("GET /login", handleLogin)
	users.HandleFunc("GET /register", handleRegisterUser)
	users.HandleFunc("GET /{userName}", handleUserProfile)
	users.HandleFunc("GET /compare", handleOverlap)

	movies := http.NewServeMux()
	movies.HandleFunc("GET /{movieName}", app.handleSingleMovie)
	movies.HandleFunc("GET /{$}", app.handleAllMovies)

	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
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

type application struct {
	queries  *database.Queries
	platform string
}
