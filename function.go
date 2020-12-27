package main

import "fmt"

func add(x int, y int) int {
	var out = x + y
	return out
}

func main() {

	num1 := 11
	num2 := 22

	result := add(num1, num2)
	fmt.Println(result)

}
