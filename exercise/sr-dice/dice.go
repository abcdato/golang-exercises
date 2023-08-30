//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
)

func rollGenerator(side int) int {
	return rand.Intn(side) + 1
}

func main() {
	dice := 2
	sides := 6
	rolls := 1

	for roll := 1; roll <= rolls; roll++ {
		sum := 0

		for die := 1; die <= dice; die++ {
			rolled := rollGenerator(sides)
			sum += rolled
			fmt.Printf("%d die rolled: %d\n", die, rolled)
		}

		fmt.Println("Totall rolled:", sum)

		switch s := sum; {
		case s == 2 && dice == 2:
			fmt.Println("Snake eyes")
		case s == 7:
			fmt.Println("Lucky 7")
		case s%2 == 0:
			fmt.Println("Even")
		case s%2 != 0:
			fmt.Println("Odd")
		}
	}

	// dice1 := roll()
	// dice2 := roll()

	// totallRoll := dice1 + dice2

	// if dice1 == 1 && dice2 == 1 {
	// 	fmt.Println("Snake eyes")
	// } else if totallRoll == 7 {
	// 	fmt.Println("Lucky 7")
	// } else if totallRoll%2 == 0 {
	// 	fmt.Println("Even")
	// } else if totallRoll%2 != 0 {
	// 	fmt.Println("Odd")
	// // }

	// fmt.Println(dice1)
	// fmt.Println(dice2)
	// fmt.Println("Dice sum is", dice1+dice2)
}
