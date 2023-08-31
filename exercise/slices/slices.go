//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part string

// * Create a function to print out the contents of the assembly line
func printParts(slice []Part) {
	for i := 0; i < len(slice); i++ {
		el := slice[i]
		fmt.Println(el)
	}
}

func main() {
	//* Using a slice, create an assembly line that contains type Part
	assemblyLine := make([]Part, 3)

	//* Perform the following:
	//  - Create an assembly line having any three parts
	fmt.Println("Created 3 parts:")
	assemblyLine[0] = "Display"
	assemblyLine[1] = "Port"
	assemblyLine[2] = "Camera"
	printParts(assemblyLine)

	//  - Add two new parts to the line
	fmt.Println("Added two new parts:")
	assemblyLine = append(assemblyLine, "Scanner", "Headphone Jack")
	printParts(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]

	//  - Print out the contents of the assembly line at each step
	fmt.Println("Sliced:")
	printParts(assemblyLine)
}
