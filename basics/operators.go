package main

import "fmt"

func main() {
	// arithmetic operators
	//examples of arithemtics :
	// + - / * %
	var num1 = 10
	var num2 = 8
	var sum = num1 + num2
	var minus = num1 - num2
	var divide = num1 / num2
	var multiplicaion = num1 * num2

	var remainder = num1 % num2

	fmt.Println(sum, minus, divide, multiplicaion, remainder)

	// relational operators

	var result = num1 > num2

	fmt.Println(result)

}
