package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func Distance(a *Point, b *Point) float64 {
	distanceSquared := math.Pow(a.GetX()-b.GetX(), 2) + math.Pow(a.GetY()-b.GetY(), 2)
	return math.Sqrt(distanceSquared)
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, -1)

	// sqrt(2)
	fmt.Println(Distance(a, b))
}
