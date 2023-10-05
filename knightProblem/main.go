package main

import (
	"errors"
	"fmt"
)

func validateInput(input string) error {
	if len(input) != 2 {
		return errors.New("input length must be 2")
	}

	if input[0] < 'a' || input[0] > 'h' {
		return errors.New("letter is not between 'a' and 'h'")
	}

	if (input[1]-'0') < 1 || (input[1]-'0') > 8 {
		return errors.New("number is not between 1 and 8")
	}

	return nil
}

func main() {
	var startPosition, endPosition string
	fmt.Scan(&startPosition)
	fmt.Scan(&endPosition)

	if err := validateInput(startPosition); err != nil {
		fmt.Println("[!! ERROR !!] | ", err)
		return
	}
	if err := validateInput(endPosition); err != nil {
		fmt.Println("[!! ERROR !!] | ", err)
		return
	}

	fmt.Printf("[startPosition] | Letter: %v | Number: %v | %v\n", string(startPosition[0]), string(startPosition[1:]), startPosition)
	fmt.Printf("[endPosition]   | Letter: %v | Number: %v | %v\n", string(endPosition[0]), string(endPosition[1:]), endPosition)
}
