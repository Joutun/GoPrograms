package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BASIC COMMANDS:
// go build first.go
// ./first.exe

// go mod init testMod

// go run first.go

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

	fmt.Println("==========")

	// print info and something variables
	fmt.Printf("info: %v\n", info)
	fmt.Printf("something: %v\n", something)

	// use function to change struct attributes
	something.changeValue(200)

	fmt.Println("==========")

	// print again to see changed values
	fmt.Println("something: ", something)

	fmt.Println("==========")
	// declare variable and take in user input
	fmt.Println("Give me a number: ")
	var number int
	amount, err := fmt.Scanf("%d", &number)

	if err == nil {

		// use user input to change value in struct
		something.changeValue(number)
		fmt.Println("Number: ", number, " read ", amount, "something: ", something)
		fmt.Println("==========")

	} else {
		fmt.Println("Error!: ", err)

		fmt.Println("==========")

	}

	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')

	// amount, err = fmt.Scanf("%d", &number)

	fmt.Println("Give me a second number: ")
	text, err := reader.ReadString('\n')

	text = strings.TrimSpace(text)

	interestingValue, err := strconv.Atoi(text)

	something.changeValue(interestingValue)
	fmt.Println("Something: ", something)
}
