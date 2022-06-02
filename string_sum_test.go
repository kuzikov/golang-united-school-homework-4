package string_sum

import (
	"log"
	"testing"
)

func TestStrigSum(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
	}{
		{expr: "1+1", expected: "2"},
		{expr: "1-1", expected: "0"},
		{expr: "-1+1+2", expected: ""},
		{expr: "-4-6", expected: "-10"},
		{expr: "-10--1", expected: "-9"},
		{expr: "-1000+500", expected: "-500"},
		{expr: "23a23+2", expected: ""},
		{expr: "3+223A", expected: ""},
	}

	for _, test := range tests {
		res, err := StringSum(test.expr)
		if err != nil {
			log.Println(err.Error())
		}
		if res != test.expected {
			t.Errorf("have: %v, want:%v\n", res, test.expected)
		}
	}
}
