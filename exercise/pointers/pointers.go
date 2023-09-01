//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Inactive = false
)

type Tag bool

type Item struct {
	name string
	tag  Tag
}

// * Create functions to activate and deactivate security tags using pointers
func activate(tag *Tag) {
	*tag = Active
}
func deactivate(tag *Tag) {
	*tag = Inactive
}

// * Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item) {
	fmt.Println("Checking out items...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func main() {
	//* Perform the following:
	//  - Create at least 4 items, all with active security tags
	shirt := Item{name: "shirt", tag: Inactive}
	pants := Item{name: "pants", tag: Inactive}
	cap := Item{name: "cap", tag: Inactive}
	jacket := Item{name: "jacket", tag: Inactive}
	// fmt.Println(shirt)
	activate(&shirt.tag)
	activate(&cap.tag)
	activate(&jacket.tag)

	//  - Store them in a slice or array
	items := []Item{shirt, pants, cap, jacket}
	fmt.Println(items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[2].tag)
	fmt.Println(items)

	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	fmt.Println(items)
}
