package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

var w Wheel

func main() {

	w = Wheel{Circle{Point{8, 8}, 10}, 20}

	// equivalent to
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 10,
		},
		Spokes: 20,
	}
	fmt.Println(w)
	fmt.Printf("%v\n", w)
	fmt.Printf("%#v\n", w)

	// w.Circle.Point.X = 42 // w.X // undefine field X
	w.X = 48
	fmt.Printf("%#v\n", w)
}
