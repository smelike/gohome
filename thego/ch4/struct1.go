package main

import (
	"fmt"
	"math"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

var employeeOfTheMonth *Employee = &dilbert

func main() {
	fmt.Println(EmployeeByID(12))
}

// syntax error why??
// (*employeeOfTheMonth).Position += " (proactive team player)"
func EmployeeByID(id int) *Employee {
	/*
		how to return an entity of Employee by pointer?
	*/
	e := new(Employee)
	// fmt.Println(e)
	*e = Employee{
		math.MaxInt,
		"Jmaes",
		"Wall Cold",
		time.Now(),
		"Programmmer",
		10000,
		math.MaxInt16,
	}
	return e
}
