package main

import (
	"fmt"
	"time"
)

// small struct
type Point struct {
	X, Y int
}

// large struct or larger struct
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	fmt.Println(Scale(Point{1, 2}, 5))

	pp := new(Point) // 返回 Pointer: *type
	*pp = Point{1, 2}
	fmt.Println(pp)  // pp -> &{1,2 }
	fmt.Println(*pp) // *pp -> {1, 2}
}

// pass struct values as arguments to function and returned from them.

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
