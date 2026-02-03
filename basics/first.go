package startingGoJourney

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
	fmt.Println("==========")

	// declare and assign.
	info := 0
	info = info + 1
	// use struct
	something := Knowledge{value: 5, identifier: "important"}

	// print info and something variables
	fmt.Printf("info: %v\n", info)
	fmt.Printf("something: %v\n", something)
	fmt.Println("==========")

	// use function to change struct attributes
	something.changeValue(200)
	// print again to see changed values
	fmt.Println("something: ", something)
	fmt.Println("==========")

	// declare variable and take in user input
	fmt.Println("Give me a number: ")
	var number int
	amount, err := fmt.Scanf("%d", &number)

	if amount == 1 {
		// use user input to change value in struct
		something.changeValue(number)
		fmt.Println("Number: ", number, " read ", amount, "something: ", something)
		fmt.Println("==========")
	} else {
		// print errors if there were any
		fmt.Println("Error!: ", err)
		fmt.Println("==========")
	}

	reader := bufio.NewReader(os.Stdin)

	// flush out the 1st scanf stored \n
	reader.ReadString('\n')

	fmt.Println("Give me a second number: ")
	// take user input and stop at the first \n. can replace with other chars e.g 'x'
	text, err := reader.ReadString('\n')

	// Remove the Enter \n and spaces
	text = strings.TrimSpace(text)

	// strconv.Atoi only takes numbers. Converts string into numbers.
	interestingValue, err := strconv.Atoi(text)

	if err == nil {
		// update struct value if user input was a valid number
		something.changeValue(interestingValue)
		fmt.Println("Something: ", something)

	} else {
		// print error
		fmt.Println("Error: ", err)
	}
}
