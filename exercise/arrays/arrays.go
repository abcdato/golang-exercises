//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type ShoppingList struct {
	price int
	name  string
}

// - The total number of items
func totalNum(list [4]ShoppingList) int {
	total := 0
	for i := 0; i < len(list); i++ {
		if list[i].name != "" {
			total += 1
		}
	}
	return total
}

// - The total cost of the items
func totalCost(list [4]ShoppingList) int {
	sum := 0
	for i := 0; i < len(list); i++ {
		sum += list[i].price
	}
	return sum
}

func main() {
	//Using an array, create a shopping list with enough room
	//  for 4 products
	//* Insert 3 products into the array
	list := [4]ShoppingList{
		{price: 100, name: "Apple"},
		{price: 50, name: "Milk"},
		{price: 20, name: "Bread"},
	}

	//* Print to the terminal:
	//  - The last item on the list
	fmt.Println("Last item on the list is", list[totalNum(list)-1].name)
	//  - The total number of items
	fmt.Println("Total number of items is", totalNum(list))
	//  - The total cost of the items
	fmt.Println("Total cost of the items is", totalCost(list))

	//* Add a fourth product to the list and print out the
	list[3] = ShoppingList{price: 150, name: "Jam"}
	//  information again
	fmt.Println("Total number of items is", len(list))
	fmt.Println("Total cost of the items is", totalCost(list))

}
