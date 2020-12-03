package calculator_test

import (
	"calculator"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
	tname string
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{ a: 2, b: 2, want: 4, tname: "Multiply two positive numbers" },
		{ a: 3, b: -2, want: -6, tname: "Multiply a negative and positive number which multiply to a negative" },
		{ a: -2, b: -5, want: 10, tname: "Multiply two negative numbers which multiply to a positive"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nMultiply(%f, %f): want %f, got %f",
				tc.tname, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()
	
	testCases := []testCase{
		{ a: 2, b: 2, want: 4, tname: "Add two positive numbers" },
		{ a: -1, b: -1, want: -2, tname: "Add two negative numbers which sum to a positive" },
		{ a: 5, b: -1, want: 4, tname: "Add a postive and a negative which sum to a positive" },
	}
	
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nAdd(%f, %f): want %f, got %f",
				tc.tname, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{ a: 4, b: 2, want: 2, tname: "Subtract a positive from a positive which sum to a positive"},
		{ a: 2, b: 4, want: -2, tname: "Subtract a positive from a positive which sum to a negative"},
		{ a: -2, b: 1, want: -3, tname: "Subtract a positive from a negative which sum to a negative"},
		{ a: -2, b: -4, want: 2, tname: "Subtract a negative from a positive which sum to a positive"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nSubtract(%f, %f): want %f, got %f",
				tc.tname, tc.a, tc.b, tc.want, got)
		}
	}
}