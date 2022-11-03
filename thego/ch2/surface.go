package main

import (
	"fmt"
	"math"
	"net/http"
)

/*
	The program illustrates floating-point graphics computations.
	It plots a function of two variable z = f(x, y) as a wire mesh 3-D surface,
	using Scalable Vector Graphic (SVG), a standard XML notation for line drawing.

	the function sin(r)/r, where r is sqrt(x*x + y*y).
*/

// Surface computes an SVG rendering of 3-D surface funtion.

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y uint
	zscale        = height * 0.4        // pixels per z uint
	angle         = math.Pi / 6         // angle of x, y axes (=30)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30), cos(30)

func main() {

	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke:grey;fill: white;stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}
func corner(i, j int) (float64, float64) {
	// find point(x,y) at corner of cell(i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z
	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
