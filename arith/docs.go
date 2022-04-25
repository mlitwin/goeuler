// The `arith` package has lower-level entities and functions not rising to the level of an algorithm.
// This is admittedly a somewhat arbitrary distinction, but is intended to be a set of primitives for
// use in (more complicated) algorithms in the algo package.
//
// The general type used for integers is `int64`, since Project Euler generally keeps itself to problems that basically fit into a 32 bit integer.
// Using 64 bits leaves you room for to do some basic calculations allowing some intermediate results which would overflow in 32 bits.
package arith
