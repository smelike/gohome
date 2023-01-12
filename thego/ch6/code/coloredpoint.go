package main

import (
	"fmt"
	"image/color"
	"test/geometry"
)

// type Point struct{ X, Y float64 }

type ColorPoint struct {
	geometry.Point
	Color color.RGBA
}

/*
	内嵌可以使我们定义字段特别多的复杂类型，
	我们可以将字段先按小类型分组，然后定义小类型的方法，之后再把它们组合起来。
*/
func main() {
	var cp ColorPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 3
	fmt.Println(cp.Y)

	// R - Red, G - Green, B - blue, A - ?
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColorPoint{geometry.Point{1, 1}, red}
	var q = ColorPoint{geometry.Point{5, 4}, blue}
	//  p.Distance(q) // cannot use q (variable of type ColorPoint)
	// as type geometry.Point in argument to p.Distance
	fmt.Println("Point Distance:", p.Distance(q.Point)) // q.Point
}
