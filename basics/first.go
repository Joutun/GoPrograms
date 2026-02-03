package main

import "fmt"

// struct == tietue
type Knowledge struct {
	value      int
	identifier string
}

// method
func (k *Knowledge) changeValue(new int) {
	k.value = new
}

func main() {
	fmt.Println("Simple text")
	// declare and assign.
	info := 0
	info = info + 1

	// use struct
	something := Knowledge{value: 5, identifier: "important"}

	fmt.Printf("info: %v\n", info)
	fmt.Printf("something: %v\n", something)

	something.changeValue(200)

	fmt.Println("something: ", something)

}
