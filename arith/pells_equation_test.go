package arith

import "testing"

func TestPellsEquation(t *testing.T) {
	x, y := GetFundamentalPellsEquationSolution(7)
	if x != 8 || y != 3 {
		t.Fatal(7, " Wrong Pell", x, y, " expected ", 8, 3)
	}
}
