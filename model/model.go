package model

type Employee struct {
	Hometown int
	Capacity int
}

type Input struct {
	NumberOfTowns int
	OfficeTown    int
	Employees     []Employee
}
