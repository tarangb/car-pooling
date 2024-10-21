package work

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"elwood/model"
	"elwood/parser"
)

// ProcessFile coordinates parsing and solution logic for each test case
func ProcessFile(filePath string) ([]string, error) {
	numCases, testCases, err := parser.ParseFile(filePath)
	if err != nil {
		return nil, err
	}

	results := make([]string, numCases)
	for i, input := range testCases {
		result := computeCars(input)
		// Uncomment this line and comment previous line to use the max optimized method.
		// result := computeCarsMaxOptimized(input)
		results[i] = result
	}

	return results, nil
}

// computeCars calculates the number of cars needed or returns "IMPOSSIBLE"
func computeCars(input model.Input) string {
	totalAvailableCapacityByTown := make(map[int]int)
	individualCarCapacityByTown := make(map[int][]int) // Initializes the map

	// Sort employees by towns
	for _, emp := range input.Employees {
		totalAvailableCapacityByTown[emp.Hometown] = totalAvailableCapacityByTown[emp.Hometown] + emp.Capacity // {1 , 3} , {2, 0}. {1, 1} // 1 -> 4, 2 -> 0
		individualCarCapacityByTown[emp.Hometown] = append(individualCarCapacityByTown[emp.Hometown], emp.Capacity) // 1 -> [1 ,3] , 2 -> [0]
	}

	// Store number of cars needed
	result := make([]int, input.NumberOfTowns+1)

	for town, capacityPerCar := range individualCarCapacityByTown {
		if town == input.OfficeTown {
			continue
		}

		employeesInTown := len(capacityPerCar)
		if employeesInTown > totalAvailableCapacityByTown[town] {
			return "IMPOSSIBLE"
		}

		// Sort drivers by capacity in descending order. This takes O(N log N) where N is the
		// number of employees commuting from this specific town.
		sort.Sort(sort.Reverse(sort.IntSlice(capacityPerCar))) // [ 3 ,1]

		carsNeeded := 0
		for _, capacity := range capacityPerCar {
			employeesInTown -= capacity
			carsNeeded++
			if employeesInTown <= 0 {
				break
			}
		}

		result[town] = carsNeeded
	}

	return formatResult(result)
}

// computeCarsMaxOptimized calculates the number of cars needed or returns "IMPOSSIBLE"
// This is max optimized and time complexity is O(N) where 𝑁 = the number of employees
// in the town holding the most employees
func computeCarsMaxOptimized(input model.Input) string {
	totalAvailableCapacityByTown := make(map[int]int)
	individualCarCapacityByTown := make(map[int][]int) // Initializes the map

	// Sort employees by towns
	for _, emp := range input.Employees {
		totalAvailableCapacityByTown[emp.Hometown] = totalAvailableCapacityByTown[emp.Hometown] + emp.Capacity
		individualCarCapacityByTown[emp.Hometown] = append(individualCarCapacityByTown[emp.Hometown], emp.Capacity) // [ 3 ,1 , 3 , 6, 0]
	}

	// Store number of cars needed
	result := make([]int, input.NumberOfTowns+1)

	for town, capacityByCar := range individualCarCapacityByTown {
		if town == input.OfficeTown {
			continue
		}

		employeesInTown := len(capacityByCar)
		if employeesInTown > totalAvailableCapacityByTown[town] {
			return "IMPOSSIBLE"
		}

		// count number of cars per capacity. it takes O(N) to compute numCarsByCapacity
		numCarsByCapacity := make(map[int]int)
		for _, carCapacity := range capacityByCar {
			numCarsByCapacity[carCapacity] = numCarsByCapacity[carCapacity] + 1
		} // {3 -> 2, 1 -> 1, 6 -> 1, 0 -> 1}

		carsNeeded := 0

		// This works and will only take O(1) since we are given that max vechicle capacity is 6.
		// If max vechicle capacity can be any aribtrary high number than computeCars will be better.
		for capacity := 6; capacity > 0; capacity-- {
			availableCars := int(math.Min(
				(math.Ceil(float64(employeesInTown) / float64(capacity))), float64(numCarsByCapacity[capacity])))
			carsNeeded += availableCars
			employeesInTown -= availableCars * capacity

			if employeesInTown <= 0 {
				break
			}

		}
		result[town] = carsNeeded
	}

	return formatResult(result)
}

// Helper to format the result for the test case output
func formatResult(result []int) string {
	var formattedResult string
	for _, r := range result[1:] { // Skip index 0 as we start from Town 1
		formattedResult += fmt.Sprintf("%d ", r)
	}

	// Remove the trailing space
	return strings.TrimSpace(formattedResult)
}
