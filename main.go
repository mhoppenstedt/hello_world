package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(math.Cos(math.Pi))
	fmt.Printf("Rectangle area: %d\n", rectArea(5, 5))
	fmt.Printf("Triangle area : %f\n", triaArea(3, 3))
}

func rectArea(x, y int) int {
	return x * y
}

func triaArea(x, y int) float32 {
	return (0.5 * float32(x*y))
}
