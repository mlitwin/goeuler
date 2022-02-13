# goeuler

* Learn go
* Learn algorithms
* Play with project Euler

This is a library of general algorithms. Mostly for use in Project Euler problems. But they are all general - no spoilers here.

## Packages

### arith

The `arith` package will have lower-level entities and functions. The general type used for integers is `int64`, since Project Euler generally keeps itself to problems that basically fit into a 32 bit integer, so 64 bits leaves you room for to do some basic calculations allowing some intermediate results which would overflow in 32 bits.


#### Iterator

* Basic abstract iterator with a `HasValue()` / `NextValue()` interface
* Supports a generic `Reduce()` method

Implementations:
* `Divisors` iterates through the divisors of n

#### Misc Types and Functions

* `C(n,k)` n choose k
* [`Factoradic`](https://en.wikipedia.org/wiki/Factorial_number_system) number, supporting conversion to a permutation.
* `RationalFraction` supporting extraction of arbitrary base `NextMantissaDigit()`
* `IsPrime()`
* `Digits(n,base)` - return digits in `base` as a slice
* `ValueOfDigits(slice,base)` - convert digit slice back to `int64`
* `IntSolveQuadradic(a,b,c)` returns integer roots of `ax^2+bx+c` as a slice, largest root first
* `InverseModN(a,n)` - modular inverse (or 0 if no inverse)
* `Integer[V]` - Interface for Integer like types. `V` is the actual type of the values, the `Integer[V]` has methods to add/subtract/multiple/divide/etc `V`'s
    * `NewIntModM(m int64)` an `Integer[int64]` for arithmetic Mod m
    * `PowOf[V any](f Integer[V], x V, n int64) V`

### algo

More complicated algorithms.

#### Heap

A min `Heap[V any, P Numeric] `, supporting a `Decrease()` operation

#### A Star

A `MinPathAStar[V any, ID comparable, W Numeric](g AStarGraph[V,ID,W], start *V, end *V) (W, []*V) ` function. Takes an `AStarGraph`, start and end vertex, returns the min weight, and the path.

An interesting design question here is how to handle the auxiliary data the algorithm needs to store about each vertex. Here we require the `AStarGraph` interface to be able to give a comparable `ID` for each vertex, so the algorithm can use that as a key to an (internal) map.

Another way to go would be to require the `AStarGraph` to be able to store (and produce) the auxiliary data itself. It seemed like most implementation would end up with some kind of map anyway, which is why I didn't go this route. 

## Admin notes

## Setup

We will use `go1.18` beta, to play with the latest features, especially generics.

Maybe use a workspace: 
https://go.googlesource.com/proposal/+/master/design/45713-workspace.md

Snippets:

* `go1.18beta1` - workspaces / generics' `alias godev=go1.18beta1`
* godev work init arith algo
* godev mod init github.com/mlitwin/goeuler/arith