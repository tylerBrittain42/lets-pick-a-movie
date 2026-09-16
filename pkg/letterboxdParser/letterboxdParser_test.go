package letterboxdParser

import (
	"slices"
	"testing"
)

const testDir = "samples"

func TestGetMovieList(t *testing.T) {
	const csvName = "single.csv"
	const testFile = testDir + "/" + csvName

	expected := []movie{{name: "The Man Who Wasn't There", year: 2001, url: "n/a"}}
	actual, err := GetListFromCSV(testFile)

	if err != nil {
		t.Errorf("unexpected error %v encountered", err)
	}

	if !slices.Equal(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}

}

// cases
// single
// muliple
// no file

// add get single movies and convert this and bove to table

func TestHasHeaders(t *testing.T) {
	type testCase struct {
		description     string
		input           []string
		expectedOuttput bool
	}
	tests := []testCase{
		{
			description:     "has header",
			input:           []string{"Date", "Name", "Year", "Letterboxd URI"},
			expectedOuttput: true,
		},
		{
			description:     "does not have header",
			input:           []string{"2025-02-09", "Wildflower", "2022", "https://boxd.it/yt94"},
			expectedOuttput: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := hasHeaders(tt.input)
			if actual != tt.expectedOuttput {
				t.Errorf("Recieved %v for line %v", actual, tt.input)
			}
		})
	}

}

func TestIsValid(t *testing.T) {
	type testCase struct {
		description     string
		input           []string
		expectedOuttput bool
	}
	tests := []testCase{
		{
			description:     "is too short",
			input:           []string{"Date", "Name", "Letterboxd URI"},
			expectedOuttput: false,
		},
		{
			description:     "Does not parse year",
			input:           []string{"2025-02-09", "Wildflower", "2a22", "https://boxd.it/yt94"},
			expectedOuttput: false,
		},
		{
			description:     "valid",
			input:           []string{"2025-02-09", "Wildflower", "2022", "https://boxd.it/yt94"},
			expectedOuttput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := isValid(tt.input)
			if actual != tt.expectedOuttput {
				t.Errorf("%v: %v marked %v", tt.description, tt.input, actual)
			}
		})
	}

}
