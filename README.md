# goeuler

* Learn go
* Learn algorithms
* Play with project Euler

This is a library of general algorithms. Mostly for use in Project Euler problems. But they are all general - no spoilers here.

## Packages

* The `arith` package will have lower-level entities and functions. See [arith.md](./docs/arith.md)
* The `algo` package complicated algorithms using `arith` for primitives. See [algo.md](./docs/algo.md)
* The `textutil` package contains for reading line based files. An maybe other stuff if it comes up. See    [textutil.md](./docs/textutil.md)

# Notes and Lessons

## Generics rule of thumb

Split the `value` type from namespace/worker type. There must be a more accepted pattern name for this. But for example, `Integer[V]` - `Integer` has the methods that act on the value type `V`. So for integers mod n, you can have `IntegerModN[int64]`: An `IntegerModN[int64]` instance acts on `int64`'s.

## Return a destructor for Defer()al

```go
scanner, close := NewFileScanner("p054_poker.txt")
defer close()
```

## Things not covered

I don't think the Project Euler corpus will exercise golang serialization / deserialization and reflection. Also probably not concurrency (though maybe some sophisticated problem will be cleaner with it).



