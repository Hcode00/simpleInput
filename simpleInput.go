package simpleInput

import (
	"bufio"
	"fmt"
	"os"
)

// input function takes in a variadic parameter of strings and returns a slice of strings.
// It prompts the user to input values for each string in the parameter and stores them in the slice.
// If no values are inputted, it returns nil.
func input(str ...string) []string {

	var inputs []string
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for _, s := range str {
		fmt.Print(s)
		scanner.Scan()
		inputs = append(inputs, scanner.Text())
	}
	if len(inputs) == 0 {
		return nil
	}
	return inputs
}
