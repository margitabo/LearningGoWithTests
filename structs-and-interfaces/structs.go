package structs_and_interfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base  float64
	Heigh float64
}

//In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.
type Shape interface {
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, c.Radius)
}

func (t Triangle) Area() float64 {
	return t.Heigh * t.Base / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Base + t.Base
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}
