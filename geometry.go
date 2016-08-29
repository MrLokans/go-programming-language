package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func (p Point) Distance(anotherPoint Point) float64 {
	return math.Hypot(anotherPoint.X-p.X, anotherPoint.Y-p.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0

	for i := range p {
		if i > 0 {
			sum += p[i-1].Distance(p[i])
		}
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{30, 99}

	path := Path{
		p,
		q,
		{21, 33},
		{1, 1},
	}

	fmt.Println(p.Distance(q))
	fmt.Println(path.Distance())
}
