package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64 = 9

	var result = math.Sqrt(number)
	var intResult = math.Round(result) // we can use Ceil, Floor instead of Round
	fmt.Printf("This is the value %.2f Thank You", intResult)
}
