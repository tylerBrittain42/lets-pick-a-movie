package letterboxdParser

import (
	"encoding/csv"
	"fmt"
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

func getMovieList() movie {
	return movie{}
}

func getMovie(line []string) movie {

	return movie{}
}

func GetListFromCSV(path string) ([]movie, error) {
	f, err := os.Open(path)
	if err != nil {
		return []movie{}, err
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return []movie{}, err
	}

	for _, record := range records {
		fmt.Println(record)
	}

	_ = path
	return []movie{}, err

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
