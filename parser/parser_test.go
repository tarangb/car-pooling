package parser

import (
	"os"
	"reflect"
	"testing"

	"car-pooling/model"
)

// Helper function to create a temporary file with content
func createTempFile(t *testing.T, content string) string {
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	return tmpFile.Name()
}

func TestParseFile(t *testing.T) {
	content := `2
				3 2
				3
				1 5
				2 3
				3 4
				5 4
				2
				1 4
				2 3`

	// Create a temp file with the content
	filePath := createTempFile(t, content)
	defer os.Remove(filePath) // Clean up the temp file after test

	// Call the ParseFile function
	testCases, inputs, err := ParseFile(filePath)

	// Expected values
	expectedInputs := []model.Input{
		{
			NumberOfTowns: 3,
			OfficeTown:    2,
			Employees: []model.Employee{
				{Hometown: 1, Capacity: 5},
				{Hometown: 2, Capacity: 3},
				{Hometown: 3, Capacity: 4},
			},
		},
		{
			NumberOfTowns: 5,
			OfficeTown:    4,
			Employees: []model.Employee{
				{Hometown: 1, Capacity: 4},
				{Hometown: 2, Capacity: 3},
			},
		},
	}

	// Validate no error occurred
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	// Validate number of test cases
	if testCases != 2 {
		t.Errorf("Expected test cases to be 2, got: %d", testCases)
	}

	// Validate the input data (use reflect.DeepEqual to compare slices and structs)
	if !reflect.DeepEqual(inputs, expectedInputs) {
		t.Errorf("Expected inputs to be %v, got: %v", expectedInputs, inputs)
	}
}

func TestParseFile_InvalidData(t *testing.T) {
	tests := []struct {
		name    string
		content string
	}{
		{
			name: "Invalid Number of Test Cases",
			content: `invalid
					  3 2
					  2
					  1 5
					  2 3`,
		},
		{
			name: "Invalid Input Format (missing officeTown or town count)",
			content: `1
					  3`,
		},
		{
			name: "Invalid Employee Data (invalid capacity and hometown)",
			content: `1
					  3 2
					  2
					  1 x
					  4 5`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temp file with the test content
			filePath := createTempFile(t, tt.content)
			defer os.Remove(filePath)

			// Call the ParseFile function
			_, _, err := ParseFile(filePath)

			// Expect an error for all invalid cases
			if err == nil {
				t.Errorf("Expected error for %s, but got nil", tt.name)
			}
		})
	}
}

func TestParseFile_FileNotFound(t *testing.T) {
	_, _, err := ParseFile("nonexistent_file.txt")
	if err == nil {
		t.Errorf("Expected an error for non-existent file, got nil")
	}
}
