package main

type Circle struct {
	X, Y, Radius int
}

type Wheel struct {
	X, Y, Radius, Spokes int
}

/*
A circle has fields for the X and Y coordinates of its center, and a Radius.
A Wheel has all the features of a Circle, plus Spokes, the number of inscribed radial spokes.

*/

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20

}
