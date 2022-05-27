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
func StringSum(input string) (output string, err error) {

	if err := checkInput(input); err != nil {
		return "", fmt.Errorf("cannot to evaluate expression: %w", err)
	}
	if len(input) == 0 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}
	if len(input) < 3 {
		return "", fmt.Errorf("not enough ops: %w", errorNotTwoOperands)
	}

	input = clear(input)

	op1, op2, sig := ops(input)

	if sig == '+' {
		return strconv.Itoa(op1 + op2), nil
	}
	return strconv.Itoa(op1 - op2), nil

}

func clear(input string) string {
	cleared := make([]rune, 0, len(input))
	for _, v := range input {
		if v != ' ' {
			cleared = append(cleared, v)
		}
	}
	return string(cleared)
}

func checkInput(input string) error {
	for _, v := range input {

		if !unicode.IsDigit(v) || v != '+' && v != '-' {
			return nil
		}
		return errorRestrictedCharSet

	}
	return nil
}

func ops(expr string) (op1, op2 int, opcode byte) {
	// op1,op2 sign
	sg1, sg2 := 1, 1
	opcode = ' '
	if expr[0] == byte('-') {
		sg1 = -1
		expr = expr[1:]
	}
	// fetch op1
	for i, ch := range expr {
		if unicode.IsDigit(ch) {
			op1 = op1*10 + int(ch-'0')
			continue
		}

		switch ch {
		case '+':
			opcode = '+'
		default:
			opcode = '-'
		}
		expr = expr[i+1:]
		break

	}

	for _, ch := range expr {
		if unicode.IsDigit(ch) {
			op2 = op2*10 + int(ch-'0')
			continue
		}
		switch ch {
		case '-':
			sg2 = -1
		default:
			sg2 = 1
		}
	}

	return op1 * sg1, op2 * sg2, opcode
}
