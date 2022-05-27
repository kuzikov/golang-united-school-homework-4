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

	if err := CheckInput(input); err != nil {
		return "", fmt.Errorf("cannot to evaluate expression: %w", err)
	}
	if len(input) == 0 {
		return "", fmt.Errorf("empty input: %w", errorEmptyInput)
	}
	if len(input) < 3 {
		return "", fmt.Errorf("not enough ops: %w", errorNotTwoOperands)
	}

	input = Clear(input)

	op1, op2, sig, err := Ops(input)
	if err != nil {
		return "", errorNotTwoOperands
	}

	if sig == '+' {
		return strconv.Itoa(op1 + op2), nil
	}
	return strconv.Itoa(op1 - op2), nil

}

// Clear ...
func Clear(input string) string {
	cleared := make([]rune, 0, len(input))
	for _, v := range input {
		if v != ' ' {
			cleared = append(cleared, v)
		}
	}
	return string(cleared)
}

// CheckInput ...
func CheckInput(input string) error {
	for _, v := range input {

		if !unicode.IsDigit(v) || v != '+' && v != '-' {
			return nil
		}
		return errorRestrictedCharSet

	}
	return nil
}

// Ops ...
func Ops(expr string) (op1, op2 int, opcode rune, err error) {
	// log.Println(expr)
	// op1,op2 sign
	sg1, sg2 := 1, 1
	opcode = ' '
	// "-4--6"
	if expr[0] == '-' {
		sg1 = -1
		expr = expr[1:]
		// log.Println("sig:-\t", expr)
	} else if expr[0] == '+' {
		sg1 = 1
		expr = expr[1:]
		// log.Println("sig1:+\t", expr)
	}

	// var o1 int
	// fmt.Scan(&o1)
	// "4--6"
	// op1, opcode
	for i, ch := range expr {
		if unicode.IsDigit(ch) {
			op1 = op1*10 + int(ch-'0')
		}
		if ch == '-' || ch == '+' && opcode == ' ' {
			opcode = ch
			expr = expr[i+1:]
			break
		}
		// "-6"
	}
	// log.Printf("num1: %v\texpr: %v\n", op1, expr)
	// fmt.Scan(&o1)

	// "-6"
	if expr[0] == '-' {
		sg2 = -1
		expr = expr[1:]
		// log.Printf("sign2: %v\texpr:%v", sg2, expr)
	} else if expr[0] == '+' {
		sg2 = 1
		expr = expr[1:]
		// log.Printf("sign2: %v\texpr:%v", sg2, expr)
	}

	// log.Println(expr)
	// fmt.Scan(&o1)

	for _, ch := range expr {
		if unicode.IsDigit(ch) {
			op2 = op2*10 + int(ch-'0')
		} else {
			return 0, 0, opcode, errorRestrictedCharSet
		}
	}

	// log.Println(expr)
	// fmt.Scan(&o1)
	// log.Printf("num1: %v oper: %q num2: %v\n", op1*sg1, opcode, op2*sg2)
	return op1 * sg1, op2 * sg2, opcode, nil
}
