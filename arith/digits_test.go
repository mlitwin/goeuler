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
			d_a := d.Digits(a)
			d_b := d.Digits(b)

			x := []int64{}
			var v_x int64

			d.Sum(&x, d_a, d_b)
			v_x = d.ValueOfDigits(x)
			if v_x != a+b {
				t.Fatal("Sum failed", a, b, x)

			}

			d.Diff(&x, d_a, d_b)
			v_x = d.ValueOfDigits(x)
			if v_x != a-b {
				t.Fatal("Diff failed", a, b, x)

			}

			d.Mul(&x, d_a, d_b)
			v_x = d.ValueOfDigits(x)
			if v_x != a*b {
				t.Fatal("Mul failed", a, b, x)
			}

			cmp := d.Cmp(d_a, d_b)
			var expected_Cmp int

			if a < b {
				expected_Cmp = -1
			} else if a > b {
				expected_Cmp = 1
			}
			if cmp != expected_Cmp {
				t.Fatal("Cmp failed", a, b, cmp)
			}

		}
	}
}
