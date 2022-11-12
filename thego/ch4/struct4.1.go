package main

/*
As the set of shapes grows, we're bound to notice similarities
and repetition amongthem,
so it may be convenient to factor out their common parts.

[factor out, 分解出]

*/

// factor: point
type Point struct {
	X, Y int
}

// factor: point and radius
type Circle struct {
	Point
	Radius int
}

// factor: point, radius and spokes
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Point.X = 8 // equivalent to w.X = 8
	w.Circle.Point.Y = 8 // equivalent to w.Y = 8
	w.Spokes = 20
}
