package algo

import (
	"fmt"
	"testing"
)

type cycleSim struct {
	cur    int64
	start  int64
	period int64
}

func (c cycleSim) Clone() CycleNode {
	q := c

	return &q
}

func (c cycleSim) Equals(iq interface{}) bool {
	q := iq.(*cycleSim)
	if c.cur < c.start || q.cur < q.start {
		return false
	}
	d := q.cur - c.cur

	return d%c.period == 0
}

func (c *cycleSim) Advance() {
	c.cur++
}

func (c *cycleSim) Assign(iq interface{}) {
	q := iq.(*cycleSim)
	c.cur = q.cur
}

func TestCycleDetection(t *testing.T) {

	c := cycleSim{0, 50, 9}

	d := NewCycleDetector(&c)
	d.FindCycleFlyod()
	fmt.Println(d)
	if d.Distance%9 != 0 {
		t.Fatal("no cycle found for", c, d.Distance)
	}

	d.FindMinCycle()

	if d.Distance != 9 {
		t.Fatal("no min cycle found for", c, d.Distance)
	}

}
