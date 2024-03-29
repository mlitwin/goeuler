package arith

// Get the fundamental solution to x^2 - D y^2 = 1, Pell's equation
//
// https://en.wikipedia.org/wiki/Pell%27s_equation
// Given a fundamental solution x=x0, y=y0, iterate to find the rest via
//  x, y = (x0*x + 3*y0*y), (x0*y + y0*x)
func GetFundamentalPellsEquationSolution(D int64) (int64, int64) {
	var p []RationalFraction
	cf, v := NewRationalSurd(D)

	for {
		a := cf.NextCFTerm(v)
		p = cf.NextCFConvergent(p, a)
		conv := p[len(p)-1]
		x, y := conv.A, conv.B
		if x*x-D*y*y == 1 {
			return x, y
		}
	}
}
