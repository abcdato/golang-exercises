//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Println("Hello,", name)
}

func msg(msg string) string {
	return msg
}

func addThreeNums(a, b, c int) int {
	return a + b + c
}

func returnAnyNum(num int) int {
	return num
}

func returnAnyTwoNums(num1, num2 int) (int, int) {
	return num1, num2
}

func main() {
	greet("Josh")

	fmt.Println(msg("Messege from function!"))

	fmt.Println(addThreeNums(3, 3, 3))

	num1, num2 := returnAnyTwoNums(1, 100)
	fmt.Println(num1, num2)
	fmt.Println(addThreeNums(returnAnyNum(3), num1, num2))

	num := returnAnyNum(5)
	fmt.Println(num)
}
