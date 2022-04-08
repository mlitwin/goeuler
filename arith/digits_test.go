package arith

import (
	"testing"
)

func TestDigitList(t *testing.T) {
	// 		t.Fatal("Falied to Let", 3, m)
	var a, b int64
	d := NewDigitList(10)
	for a = 1; a < 100; a++ {
		for b = 1; b < 100; b++ {
			d_a := Digits(a, 10)
			d_b := Digits(b, 10)

			x := []int64{}
			var v_x int64

			d.Sum(&x, d_a, d_b)
			v_x = ValueOfDigits(x, 10)
			if v_x != a+b {
				t.Fatal("Sum failed", a, b, x)

			}

			d.Diff(&x, d_a, d_b)
			v_x = ValueOfDigits(x, 10)
			if v_x != a-b {
				t.Fatal("Diff failed", a, b, x)

			}

			d.Mul(&x, d_a, d_b)
			v_x = ValueOfDigits(x, 10)
			if v_x != a*b {
				t.Fatal("Mul failed", a, b, x)

			}

		}
	}
}
