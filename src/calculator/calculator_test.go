package calculator_test

import (
	"calculator"
	"testing"
	"math"
	"math/rand"
	"time"
)

type testCase struct {
	a, b float64
	want float64
	name string
}

type testCaseDivide struct {
	a, b        float64
	want        float64
	errExpected bool
	name        string
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Multiply two positive numbers"},
		{a: 3, b: -2, want: -6, name: "Multiply a negative and positive number which multiply to a negative"},
		{a: -2, b: -5, want: 10, name: "Multiply two negative numbers which multiply to a positive"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nMultiply(%f, %f): want %f, got %f",
				tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Add two positive numbers"},
		{a: -1, b: -1, want: -2, name: "Add two negative numbers which sum to a positive"},
		{a: 5, b: -1, want: 4, name: "Add a postive and a negative which sum to a positive"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nAdd(%f, %f): want %f, got %f",
				tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		// if want != got {
		if !closeEnough(want, got, 0.00000001) {
			t.Fatalf("Add(%f, %f): want %f, got %f",
				a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "Subtract a positive from a positive which sum to a positive"},
		{a: 2, b: 4, want: -2, name: "Subtract a positive from a positive which sum to a negative"},
		{a: -2, b: 1, want: -3, name: "Subtract a positive from a negative which sum to a negative"},
		{a: -2, b: -4, want: 2, name: "Subtract a negative from a positive which sum to a positive"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s\nSubtract(%f, %f): want %f, got %f",
				tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testCases := []testCaseDivide{
		{a: 4, b: 2, want: 2, errExpected: false, name: "Divide a positive by a smaller positive which result to a positive"},
		{a: 4, b: 8, want: 0.5, errExpected: false, name: "Divide a positive by a larger positive which result to a positive"},
		{a: 4, b: 0, want: 0, errExpected: true, name: "Error, not able to divide by zero"},
		{a: 4, b: 0, want: 999, errExpected: true, name: "Error, not able to divide by zero, but result returned not zero"},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errRecieved := err != nil

		if tc.errExpected != errRecieved {
			t.Fatalf("%s\nDivide(%f, %f): unexpected error status: %v",
				tc.name, tc.a, tc.b, errRecieved)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s\nDivide(%f, %f): want %f, got %f",
				tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}