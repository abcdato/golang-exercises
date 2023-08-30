//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x, y int
}

// * Create a rectangle structure containing its coordinates
type Rectangle struct {
	a Coordinate // top left
	b Coordinate // bottom right
}

func width(rect Rectangle) int {
	return rect.b.x - rect.a.x
}

func lenght(rect Rectangle) int {
	return rect.a.y - rect.b.y
}

func area(rect Rectangle) int {
	return width(rect) * lenght(rect)
}

func perimeter(rect Rectangle) int {
	return (width(rect) + lenght(rect)) * 2
}

// * Using functions, calculate the area and perimeter of a rectangle,
//   - Print the results to the terminal
//   - The functions must use the rectangle structure as the function parameter
func printInfo(rect Rectangle) {
	fmt.Println("Area is", area(rect))
	fmt.Println("Perimetr is", perimeter(rect))
}

func main() {
	rect := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printInfo(rect)

	//* After performing the above requirements, double the size
	//  of the existing rectangle and repeat the calculations
	rect.a.y *= 2
	rect.b.x *= 2
	//  - Print the new results to the terminal
	printInfo(rect)
}

// func rectangleArea(a, b int) int {
// 	return a * b
// }

// func rectanglePerimeter(a, b int) int {
// 	return (a + b) * 2
// }

// func main() {
// 	type Rectangle struct {
// 		width int
// 		lenght int
// 	}

// 	rect := Rectangle{3, 4}

// 	area := rectangleArea(rect.width, rect.lenght)
// 	fmt.Println("Area of a rectangle is", area)

// 	perimeter := rectanglePerimeter(rect.width, rect.lenght)
// 	fmt.Println("Perimeter of a rectangle is", perimeter)
// }
