package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func (p Point) Distance(anotherPoint Point) float64 {
	return math.Hypot(anotherPoint.X-p.X, anotherPoint.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	q := Point{30, 99}

	fmt.Println(p.Distance(q))
}
