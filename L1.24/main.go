package main

import (
	"fmt"
	"math"
)

func main() {
	p1 := newPoint(0, 0)
	p2 := newPoint(3, 4)
	dist := p1.Distance(p2)
	fmt.Println(dist)
}

// Point — структура с приватными полями x и y
type Point struct {
	x float64
	y float64
}

// NewPoint — конструктор для создания точки
func newPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance — метод для вычисления расстояния до другой точки
func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}
