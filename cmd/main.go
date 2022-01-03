package main

import (
	"fmt"
	"github.com/exxxception/learn-go/pkg"
)

func main() {
	circle1 := pkg.NewCircle(4)
	fmt.Printf("Area: %f\n", circle1.CalculateCircleArea())
}
