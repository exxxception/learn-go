package main

import (
	"fmt"
	"github.com/exxxception/learn-go/pkg"
)

func main() {
	result := pkg.ToCamelCase("Hi -My name_is VLad")
	fmt.Println(result)
}
