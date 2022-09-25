package arith

import "testing"

func TestPellsEquation(t *testing.T) {
	_, f := GetFundamentalPellsEquationSolution(7)
	expected := RationalSurdValue{8, 3, 1}
	if *f != expected {
		t.Fatal(7, " Wrong Pell", *f, " expected ", expected)
	}
}
