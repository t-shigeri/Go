package main

import  "fmt"
import "math"

type Point struct {
	X, Y int
}

func (p Point) Distance(q Point) int {
	d := math.Sqrt(
		float64((p.X-q.X)*(p.X-q.X) +
				(p.Y-q.Y)*(p.Y-q.Y)))
		return int(d)
}
p1 := Point{5,5}
p2 := Point{8,9}
fmt.Println(p1.Distance(p2))