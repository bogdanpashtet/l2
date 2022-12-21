package pattern

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea()
	accept(Vizitor)
}

type Vizitor interface {
	vizitSquare(*Square) int
	vizitCircle(*Circle) float64
	vizitTriangle(*Triangle) float64
}

type Square struct {
	side int
}

func (s *Square) getArea() int {
	return s.side * s.side
}

func (s *Square) accept(v Vizitor) int {
	return v.vizitSquare(s)
}

type Circle struct {
	radius int
}

func (c *Circle) getArea() float64 {
	return math.Pow(float64(c.radius), 2) * math.Pi
}

func (s *Circle) accept(v Vizitor) float64 {
	return v.vizitCircle(s)
}

type Triangle struct {
	side1, side2, side3 float64
}

func (t *Triangle) getArea() float64 {
	p := (t.side1 + t.side2 + t.side3) / 2
	h1 := 2 / t.side1 * math.Sqrt(p*(p-t.side1)*(p-t.side2)*(p-t.side3))
	return h1 * t.side1 / 2
}

func (s *Triangle) accept(v Vizitor) float64 {
	return v.vizitTriangle(s)
}

type Perimeter struct {
	p int
}

func (per *Perimeter) vizitSquare(s *Square) int {
	return s.side * 4
}

func (per *Perimeter) vizitCircle(c *Circle) float64 {
	return 2 * math.Pi * float64(c.radius)
}

func (per *Perimeter) vizitTriangle(t *Triangle) float64 {
	return t.side1 + t.side2 + t.side3
}

func VizitorFunc() {
	var s = Square{
		side: 5,
	}
	var c = Circle{
		radius: 3,
	}
	var t = Triangle{
		side1: 3,
		side2: 4,
		side3: 5,
	}
	fmt.Printf("Площадь квадрата со стороной %v  =  %v\n", s.side, s.getArea())
	fmt.Printf("Площадь круга с радиусом %v = %v\n", c.radius, c.getArea())
	fmt.Printf("Площадь треугольника со сторонами %v,%v,%v = %v\n", t.side1, t.side2, t.side3, t.getArea())

	p := &Perimeter{}
	fmt.Printf("Периметр квадрата со стороной %v  =  %v\n", s.side, s.accept(p))
	fmt.Printf("Периметр круга с радиусом %v = %v\n", c.radius, c.accept(p))
	fmt.Printf("Периметр треугольника со сторонами %v,%v,%v = %v\n", t.side1, t.side2, t.side3, t.accept(p))
}
