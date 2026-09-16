package letterboxdParser

import (
	"encoding/csv"
	"fmt"
	"os"
)

type movie struct {
	name string
	year int
	url  string
}

func getMovieList() movie {
	return movie{}
}

func getMovie() movie {

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

func hasHeaders(line []string) bool {
	return false
}
func isValid(line []string) bool {
	return false
}
