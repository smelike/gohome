package main

import (
	"fmt"
	"math"
	"time"
)

type Rocket struct {
	/*  */
}

func (r *Rocket) Launch() {
	fmt.Println("Launch:", time.Now())
}

type Point struct {
	X, Y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point { return Point{p.X + q.X, p.Y + q.Y} }
func (p Point) Sub(q Point) Point { return Point{p.X - q.X, p.Y - q.Y} }

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	// r := new(Rocket)
	// 传入匿名函数
	/* time.AfterFunc(2*time.Second, func() {
		r.Launch()
	}) */

	// 直接使用方法“值”传入 AfterFunc
	// time.AfterFunc(2*time.Second, r.Launch)

	// time.Sleep(10 * time.Second)

	p := Point{1, 2}
	q := Point{4, 6}
	distance := Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)

	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)

	// the go programming language 中函数和方法的区别是指有没有接收器，
	// 而不像其他语言那样是指有没有返回值。
}
