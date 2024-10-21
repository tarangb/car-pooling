package work

import (
	"testing"

	"elwood/model"
)

// TestComputeCars function for testing the computeCars function.
func TestComputeCars(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Input
		expected string
	}{
		{
			name: "Sufficient Capacity",
			input: model.Input{
				NumberOfTowns: 2,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 2},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 0},
				},
			},
			expected: "0 3",
		},
		{
			name: "Sufficient Capacity Multiple Towns",
			input: model.Input{
				NumberOfTowns: 4,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 2},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 2},
					{Hometown: 3, Capacity: 1},
					{Hometown: 3, Capacity: 1},
					{Hometown: 3, Capacity: 3},
					{Hometown: 4, Capacity: 2},
					{Hometown: 4, Capacity: 1},
					{Hometown: 4, Capacity: 1},
					{Hometown: 4, Capacity: 0},
				},
			},
			expected: "0 3 1 3",
		},
		{
			name: "Insufficient Capacity",
			input: model.Input{
				NumberOfTowns: 4,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 1},
				},
			},
			expected: "IMPOSSIBLE",
		},
		{
			name: "No Employees Case",
			input: model.Input{
				NumberOfTowns: 2,
				OfficeTown:    1,
				Employees:     []model.Employee{},
			},
			expected: "0 0",
		},
		{
			name: "Single Town Only",
			input: model.Input{
				NumberOfTowns: 1,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 1, Capacity: 0},
					{Hometown: 1, Capacity: 0},
				},
			},
			expected: "0", // Since the office is in town 1, no cars are needed
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := computeCars(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestComputeCarsMaxOptimized(t *testing.T) {
	tests := []struct {
		name     string
		input    model.Input
		expected string
	}{
		{
			name: "Sufficient Capacity",
			input: model.Input{
				NumberOfTowns: 2,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 2},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 0},
				},
			},
			expected: "0 3",
		},
		{
			name: "Sufficient Capacity Multiple Towns",
			input: model.Input{
				NumberOfTowns: 4,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 2},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 1},
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 2},
					{Hometown: 3, Capacity: 1},
					{Hometown: 3, Capacity: 1},
					{Hometown: 3, Capacity: 3},
					{Hometown: 4, Capacity: 2},
					{Hometown: 4, Capacity: 1},
					{Hometown: 4, Capacity: 1},
					{Hometown: 4, Capacity: 0},
				},
			},
			expected: "0 3 1 3",
		},
		{
			name: "Insufficient Capacity",
			input: model.Input{
				NumberOfTowns: 4,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 0},
					{Hometown: 2, Capacity: 1},
				},
			},
			expected: "IMPOSSIBLE",
		},
		{
			name: "No Employees Case",
			input: model.Input{
				NumberOfTowns: 2,
				OfficeTown:    1,
				Employees:     []model.Employee{},
			},
			expected: "0 0",
		},
		{
			name: "Single Town Only",
			input: model.Input{
				NumberOfTowns: 1,
				OfficeTown:    1,
				Employees: []model.Employee{
					{Hometown: 1, Capacity: 0},
					{Hometown: 1, Capacity: 0},
				},
			},
			expected: "0", // Since the office is in town 1, no cars are needed
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := computeCarsMaxOptimized(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
