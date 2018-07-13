package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Cos(math.Pi))
	fmt.Println("Rectangle area: ", rectArea(5, 5))
}

func rectArea(x, y int) int {
	return x * y
}
