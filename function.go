package main

import "fmt"

func calc(x int, y int) (out1, out2 int) {
	out1 = x + y
	out2 = x - y
	return
}

func main() {

	num1 := 110
	num2 := 22

	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)

}
