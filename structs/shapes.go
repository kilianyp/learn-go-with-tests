package main

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Triangle) Area() float64 {
	return r.Height * r.Width / 2
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
