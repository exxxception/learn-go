package pkg

import "math"

type Circle struct {
	radius int
}

func NewCircle(radius int) Circle {
	return Circle{radius: radius}
}

func (c Circle) CalculateCircleArea() float32 {
	return float32(c.radius) * float32(c.radius) * math.Pi
}
