package letterboxdParser

import (
	"encoding/csv"
	"os"
	"slices"
	"strconv"
)

var headerLine = []string{"Date", "Name", "Year", "Letterboxd URI"}

type movie struct {
	name string
	year int
	url  string
}

// Returns a parsed list of movies from a csv
// handles validation
// assumes it is a letterboxd csv
func GetListFromCSV(path string) ([]movie, error) {

	// get list
	f, err := os.Open(path)
	if err != nil {
		return []movie{}, err
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return []movie{}, err
	}

	if hasHeaders(records[0]) {
		records = records[1:]
	}

	movies := getMovieList(records)

	return movies, nil

}

// Parses 2d array into movie list
func getMovieList(lines [][]string) []movie {
	movies := []movie{}

	for _, mov := range lines {
		if isValid(mov) {
			movies = append(movies, getMovie(mov))
		}

	}
	return movies
}

// Returns the line fit into a movie
//
// Assumes that it has already been validated
func getMovie(line []string) movie {
	name := line[1]
	url := line[3]
	year, _ := strconv.Atoi(line[2])

	return movie{name, year, url}
}

// Returns true if given line is the letterboxd header
func hasHeaders(line []string) bool {
	return slices.Equal(line, headerLine)
}

// Returns true if the following criteria are met for a given line
// length = length of header
// the third value can be parsed to an int
//
// Note: while a header would not pass,
// this does not distinguish a header from invalid
func isValid(line []string) bool {
	if len(line) != len(headerLine) {
		return false
	}
	_, err := strconv.Atoi(line[2])
	if err != nil {
		return false
	}

	return true
}
