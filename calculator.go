package main

import "fmt"

func main() {
	fmt.Println("num1: ")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("num2: ")
	var num2 int
	fmt.Scanln(&num2)
	fmt.Println("operation: ")
	var operation string
	fmt.Scanln(&operation)

	var result int
	if operation == "+" {
		result = num1 + num2
	} else if operation == "-" {
		result = num1 - num2
	} else if operation == "*" {
		result = num1 * num2
	} else if operation == "/" {
		if num2 == 0 {
			fmt.Println("no.cannot be divided")
			return
		}
		result = num1 / num2
	}

	fmt.Println("The result is: ", result)
}
