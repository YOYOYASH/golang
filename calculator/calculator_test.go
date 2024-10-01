package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

type testCase struct {
	a, b float64
	want float64
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b,
				tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 0, want: 2},
		{a: 4, b: 9, want: -5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 0, want: 0},
		{a: 4, b: 9, want: 36},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2, b: 1, want: 2},
		{a: 4, b: 2, want: 2},
		{a:1,b:3, want: 0.33333},
		
	}
	for _, tc := range testCases {
		got, error := calculator.Divide(tc.a, tc.b)
		if error != nil {
			t.Fatalf("Invalid input received")
		}
		if !closeEnough(tc.want,got,0.01){
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 2,  want: 1.41},
		{a: 4,  want: 2},
		{a:1, want: 1},
		
	}
	for _, tc := range testCases {
		got, error := calculator.Sqrt(tc.a)
		if error != nil {
			t.Fatalf("Invalid input received")
		}
		if !closeEnough(tc.want,got,0.01){
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}

	}
}


func closeEnough(a,b,tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
} 