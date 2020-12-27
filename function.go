package main

import "fmt"

func calc(x int, y int) (int, int) {
	var out = x + y
	var out2 = x - y
	return out, out2
}

func main() {

	num1 := 110
	num2 := 22

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)

}
