package algo

type CycleNode interface {
	Advance() // ++
	Equals(c interface{}) bool
	Assign(c interface{}) // Set this to c
	Clone() CycleNode     // Return a copy

}

type CycleDetector struct {
	Start    CycleNode
	Tortice  CycleNode
	Hare     CycleNode
	Distance int64
}

func NewCycleDetector(start CycleNode) *CycleDetector {
	d := CycleDetector{start, nil, nil, 0}
	d.Tortice = start.Clone()
	d.Hare = start.Clone()

	return &d
}

// Advance to a cycle multiple
func (d *CycleDetector) FindCycleFlyod() {
	for {
		d.Tortice.Advance()
		d.Distance++
		if d.Tortice.Equals(d.Hare) {
			return
		}

		d.Tortice.Advance()
		d.Distance++

		if d.Tortice.Equals(d.Hare) {
			return
		}

		d.Hare.Advance()

		d.Distance--
	}
}

// At a cycle multiple? Find the min period
func (d *CycleDetector) FindMinCycle() int64 {
	d.Hare.Assign(d.Tortice)
	d.Distance = 0
	for {
		d.Hare.Advance()
		d.Distance++
		if d.Tortice.Equals(d.Hare) {
			return d.Distance
		}
	}
}

// At a cycle ? Find where cycling starts
func (d *CycleDetector) FindCycleStart() int64 {
	var torticeMove int64
	d.Tortice.Assign(d.Start)
	d.Hare.Assign(d.Start)
	cycleDistance := d.Distance
	d.Distance = 0
	for cycleDistance > 0 {
		d.Hare.Advance()
		d.Distance++
		cycleDistance--
	}

	for !d.Hare.Equals(d.Tortice) {
		d.Hare.Advance()
		d.Tortice.Advance()
		torticeMove++
	}

	return torticeMove
}
