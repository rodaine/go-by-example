package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

// LSP Example #0: squares don't have "width" & "height" #lspedant
type Square struct {
	size float64
}

func (s Square) area() float64 {
	return math.Pow(s.size, 2)
}

func (s Square) perim() float64 {
	return 4 * s.size
}

type Rect struct {
	width, height float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perim() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	geo := []Geometry{Square{size: 2}, Rect{width: 3, height: 4}, Circle{radius: 1}}
	for _, g := range geo {
		measure(g)
	}
}
