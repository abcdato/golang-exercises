//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favColor string = "green"
	fmt.Println("Favorite color is", favColor)

	birthYear, age := 1992, 31
	fmt.Println("My birth year is", birthYear, "and I'm", age, "is old")

	var (
		firstInitial = "D"
		lastInitial  = "K"
	)
	fmt.Println(firstInitial)
	fmt.Println(lastInitial)

	var ageInDays int
	ageInDays = age * 365
	fmt.Println("I'm", ageInDays, "days old")
}
