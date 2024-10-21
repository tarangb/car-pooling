package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"elwood/model"
)

func ParseFile(filePath string) (int, []model.Input, error) {
	log.Printf("Starting to parse file: %s", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer func() {
		log.Printf("Closing file: %s", filePath)
		file.Close()
	}()

	scanner := bufio.NewScanner(file)

	// Read number of test cases
	if !scanner.Scan() {
		log.Printf("Failed to read number of test cases from file: %s", filePath)
		return 0, nil, fmt.Errorf("failed to read number of test cases")
	}
	// Trim spaces from the input and convert to integer
	testCases, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		log.Printf("Invalid number of test cases: %s", err)
		return 0, nil, fmt.Errorf("invalid number of test cases")
	}

	// Prepare to collect all test case data
	testCasesInput := make([]model.Input, testCases)

	// Process each test case
	for testCase := 0; testCase < testCases; testCase++ {
		// Read number of towns and office town, trimming spaces
		if !scanner.Scan() {
			log.Printf("Failed to read towns and office info for test case #%d", testCase+1)
			return 0, nil, fmt.Errorf("failed to read towns and office info")
		}
		parts := strings.Fields(strings.TrimSpace(scanner.Text())) // Automatically trims and splits
		if len(parts) < 2 {
			log.Printf("Invalid input format for towns and office in test case #%d", testCase+1)
			return 0, nil, fmt.Errorf("invalid input format for towns and office")
		}

		numTowns, err1 := strconv.Atoi(parts[0])
		officeTown, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil || numTowns <= 0 || officeTown <= 0 || officeTown > numTowns {
			log.Printf("Invalid values for towns or office location in test case #%d: numTowns: %d, officeTown: %d", testCase+1, numTowns, officeTown)
			return 0, nil, fmt.Errorf("invalid values for towns or office location")
		}

		// Read number of employees
		if !scanner.Scan() {
			log.Printf("Error reading number of employees for test case #%d", testCase+1)
			return 0, nil, fmt.Errorf("failed to read number of employees")
		}
		numEmployees, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || numEmployees <= 0 {
			log.Printf("Invalid number of employees for test case #%d: %s", testCase+1, err)
			return 0, nil, fmt.Errorf("invalid number of employees")
		}

		employees := make([]model.Employee, numEmployees)
		for ii := 0; ii < numEmployees; ii++ {
			if !scanner.Scan() {
				log.Printf("Error reading employee #%d data in test case #%d", ii+1, testCase+1)
				return 0, nil, fmt.Errorf("failed to read employee data")
			}
			parts = strings.Fields(strings.TrimSpace(scanner.Text())) // Automatically trims and splits
			if len(parts) < 2 {
				log.Printf("Invalid employee data format for employee #%d in test case #%d", ii+1, testCase+1)
				return 0, nil, fmt.Errorf("invalid employee data format")
			}

			hometown, err1 := strconv.Atoi(parts[0])
			capacity, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil || hometown <= 0 || hometown > numTowns || capacity < 0 || capacity > 6 {
				log.Printf("Invalid values for employee #%d in test case #%d: hometown: %d, capacity: %d", ii+1, testCase+1, hometown, capacity)
				return 0, nil, fmt.Errorf("invalid values for employee hometown or capacity")
			}
			employees[ii] = model.Employee{Hometown: hometown, Capacity: capacity}
		}

		// Add test case employees to the overall list
		testCasesInput[testCase] = model.Input{NumberOfTowns: numTowns, OfficeTown: officeTown, Employees: employees}
	}

	log.Printf("Finished parsing file: %s", filePath)
	return testCases, testCasesInput, nil
}
