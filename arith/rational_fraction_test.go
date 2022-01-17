package arith

import (
	"testing"
)

func TestNewRationalFraction(t *testing.T) {
	f := NewRationalFraction(2, 3)
	if f.A != 2 {
		t.Fatal("Falied to construct numerator")
	}

	if f.B != 3 {
		t.Fatal("Falied to construct denominator")
	}
}

func TestNextMantissaDigit(t *testing.T) {
	f := NewRationalFraction(1, 3)

	d := f.NextMantissaDigit(10)
	if d != 3 {
		t.Fatal("Falied to get mantissa digit", f)
	}

	if f.A != 1 {
		t.Fatal("Falied to set remainder mantissa digit", f)
	}

	f = NewRationalFraction(1, 7)

	d = f.NextMantissaDigit(10)
	if d != 1 {
		t.Fatal("Falied to get mantissa digit of", f, "found", d)
	}

	d = f.NextMantissaDigit(10)
	if d != 4 {
		t.Fatal("Falied to get mantissa digit of", f, "found", d)
	}

}
