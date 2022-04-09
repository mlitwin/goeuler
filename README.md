# goeuler

* Learn go
* Learn algorithms
* Play with project Euler

This is a library of general algorithms. Mostly for use in Project Euler problems. But they are all general - no spoilers here.

## Packages

### arith

The `arith` package will have lower-level entities and functions. The general type used for integers is `int64`, since Project Euler generally keeps itself to problems that basically fit into a 32 bit integer, so 64 bits leaves you room for to do some basic calculations allowing some intermediate results which would overflow in 32 bits.

[arith.md](./docs/arith.md)

### algo

More complicated algorithms.

[algo.md](./docs/algo.md)

### textutil

Utilities for reading line based files. An maybe other stuff if it comes up.

[textutil.md](./docs/textutil.md)
