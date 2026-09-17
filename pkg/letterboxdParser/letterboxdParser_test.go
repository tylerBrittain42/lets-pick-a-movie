package letterboxdParser

import (
	"slices"
	"strings"
	"testing"
)

const testDir = "samples"

func TestGetListFromCsv(t *testing.T) {
	type testCase struct {
		description    string
		inputFile      string
		expectedOutput []movie
		shouldErr      bool
		errMessage     string
	}
	tests := []testCase{
		{
			description:    "single movie",
			inputFile:      "single.csv",
			expectedOutput: []movie{{name: "The Man Who Wasn't There", year: 2001, url: "https://boxd.it/1Vqc"}},
			shouldErr:      false,
			errMessage:     "",
		},
		{
			description: "list - with header",
			inputFile:   "list.csv",
			expectedOutput: []movie{
				{name: "The Man Who Wasn't There", year: 2001, url: "https://boxd.it/1Vqc"},
				{name: "Wildflower", year: 2022, url: "https://boxd.it/yt94"},
				{name: "1992", year: 2022, url: "https://boxd.it/eB6C"},
				{name: "The Inauguration of the Pleasure Dome", year: 1954, url: "https://boxd.it/UJU"},
			},
			shouldErr:  false,
			errMessage: "",
		},
		{
			description: "list - without header",
			inputFile:   "listNoHeader.csv",
			expectedOutput: []movie{
				{name: "The Man Who Wasn't There", year: 2001, url: "https://boxd.it/1Vqc"},
				{name: "Wildflower", year: 2022, url: "https://boxd.it/yt94"},
				{name: "1992", year: 2022, url: "https://boxd.it/eB6C"},
				{name: "The Inauguration of the Pleasure Dome", year: 1954, url: "https://boxd.it/UJU"},
			},
			shouldErr:  false,
			errMessage: "",
		},

		{
			description:    "bad file(path is bad)",
			inputFile:      "thisFileDoesNotExist.csv",
			expectedOutput: []movie{},
			shouldErr:      true,
			errMessage:     "no such file or directory",
		},

		{
			description:    "bad file(invalid)",
			inputFile:      "invalid.csv",
			expectedOutput: []movie{},

			shouldErr:  true,
			errMessage: "wrong number of fields",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			testFilePath := testDir + "/" + tt.inputFile
			actual, err := GetListFromCSV(testFilePath)

			// handling error cases
			if err != nil {
				if !tt.shouldErr {
					t.Errorf("Unexpected error %v occured", err)
				} else if !strings.Contains(err.Error(), tt.errMessage) {
					t.Errorf("Incorrect error message: Expected %v got %v", tt.errMessage, err)
				}
			}

			if !slices.Equal(actual, tt.expectedOutput) {
				t.Errorf("expected %v, got %v", tt.expectedOutput, actual)
			}

		})
	}

}

func TestGetMovie(t *testing.T) {

	type testCase struct {
		description     string
		input           []string
		expectedOuttput movie
	}
	tests := []testCase{
		{
			description:     "normal",
			input:           []string{"2025-02-09", "Wildflower", "2022", "https://boxd.it/yt94"},
			expectedOuttput: movie{name: "Wildflower", year: 2022, url: "https://boxd.it/yt94"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			actual := getMovie(tt.input)
			if actual != tt.expectedOuttput {
				t.Errorf("expected %v got %v", tt.expectedOuttput, actual)
			}
		})

	}
}

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
