## Project Structure

    main.go - main file
    model
        model.go - Model for processing input
    parser
        parser.go - Parses the  input file
        work_test.go - Tests for parsing file
    work
        work.go - Computes the cars needed
        work_test.go - Tests for computing minimum vehicles needed
    testdata
        input.txt - Test cases


## Setup

To setup the module run the following command in **elwood** directory

    go mod init elwood


## Run

You can run the project by specifying the input file as a command-line argument.

    go run main.go testdata/input.txt


## Test

To run tests, run the following command

    go test ./...


## Assumptions

* No cars is needed for transportation if employee home town is in the office town
* **computeCars** has a complexity of O(N log N) where N = the number of employees in the town holding the most employees
* **computeCarsMaxOptimized** will have a time complexity of O(N) where N = the number of employees in the town holding the most employees, if the number of vehicle capacity will not be an arbitrary high random number as compared to the total number of employees.