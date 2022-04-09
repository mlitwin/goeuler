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

Utilities for reading line based files. An maybe other stuff if it comes up

#### NewFileScanner

Create a `bufio.Scanner` from a file name

```
	scanner, close := NewFileScanner("p054_poker.txt")
	defer close()
	for scanner.Scan() {
		line := scanner.Text()
	}
```

## Admin notes

## Setup

We will use `go1.18` beta, to play with the latest features, especially generics.

Maybe use a workspace: 
https://go.googlesource.com/proposal/+/master/design/45713-workspace.md

Snippets:

* `go1.18beta1` - workspaces / generics' `alias godev=go1.18beta1`
* godev work init arith algo
* godev mod init github.com/mlitwin/goeuler/arith

### Docs

https://github.com/princjef/gomarkdoc
