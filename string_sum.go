package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
	// Use when encounter restricted character in math exppession.
	errorRestrictedCharSet = errors.New("math expression contains restricted characters")
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

// StringSum evaluate simple math expression. This is ugly peace of sheat. Got to be rewrited.
func StringSum(input string) (string, error) {

	if len(input) < 3 {
		return "", errorEmptyInput
	}
	for _, ch := range input { //+
		if !unicode.IsDigit(ch) && !(ch == '+' || ch == '-') {
			return "", errorRestrictedCharSet
		}
	}

	sigop1 := 0
	action := ' '
	switch {
	case input[0] == '-':
		sigop1 = -1
		input = input[1:]

	case input[0] == '+':
		sigop1 = 1
		input = input[1:]

	default:
		sigop1 = 1
	}

	//     i
	// "125+-25"
	tmp := ""
	ops := []string{}
	for i, ch := range input {
		// +
		if unicode.IsDigit(ch) {
			tmp = tmp + string(ch)
		} else {
			ops = append(ops, tmp)
			action = ch
			input = input[i+1:]
			tmp = ""
			break
		}
	}

	// "-25"

	// if !unicode.IsDigit(rune(input[0])) {
	// 	for _, v := range input {
	// 		if
	// 	}
	// }

	for _, ch := range input {
		tmp = tmp + string(ch)
	}
	ops = append(ops, tmp)
	tmp = ""

	num1, err1 := strconv.Atoi(ops[0])
	if err1 != nil {
		return "", fmt.Errorf("operand1 not valid: %w", err1)
	}
	num2, err2 := strconv.Atoi(ops[1])
	if err2 != nil {
		return "", fmt.Errorf("operand1 not valid: %w", err2)
	}

	if action == '+' {
		return strconv.Itoa((num1 * sigop1) + num2), nil
	} else if action == '-' {
		return strconv.Itoa((num1 * sigop1) - num2), nil
	}
	return "", fmt.Errorf("unknown input: %w", errorNotTwoOperands)
}
