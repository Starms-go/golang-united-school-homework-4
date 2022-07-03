package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	var sum int

	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}
	input = strings.TrimLeft(input, "+")
	input = strings.ReplaceAll(input, "+", "~+")
	input = strings.ReplaceAll(input, "-", "~-")
	input = strings.TrimLeft(input, "~")
	inputA := strings.Split(input, "~")

	if len(inputA) != 2 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	for _, y := range inputA {
		v, e := strconv.Atoi(y)
		err := fmt.Errorf("bad token. %w", e)
		if err != nil {
			fmt.Println(err.Error())
			return "", err
		} else {
			sum += v
		}
	}

	output = strconv.Itoa(sum)

	return
}
