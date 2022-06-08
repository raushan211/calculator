package main

import "fmt"

func main() {
	num1 := 100
	num2 := 500
	result := 0
	operation := "+"
	if operation == "+" {
		result = num1 + num2
	}
	if operation == "-" {
		result = num1 - num2
	}
	if operation == "/" {
		result = num1 / num2
	}
	if operation == "*" {
		result = num1 * num2
	}
	fmt.Println("The result is: ", result)
}
