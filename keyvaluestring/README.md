# Key/Values to String

This benchmark tests different ways to transform a set of key/value pairs to a combined string.

## Expectations

Given key/value pairs, expect a single string representation with ordered keys.

```
given = {
  foo = FOO
  bar = BAR
}

expect = "bar=BAR,foo=FOO"
```

## Example Use-Cases
- Query strings (use `Values#Encode` from `net/url` instead)
- Environment Variables for Executing a Command (`Command#Env` takes in a slice, use that instead)
- LDFlags (my use-case)

## Running the Benchmark

```
$ make bench
```

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/kevinmartin/go-benchmarks/keyvaluestring
Benchmark/MapWithSort-8                     10000000	       378 ns/op	      96 B/op	       5 allocs/op
Benchmark/MatrixStringJoin-8                20000000	       190 ns/op	      64 B/op	       4 allocs/op
Benchmark/MatrixStringBuilder-8             50000000	       107 ns/op	      24 B/op	       2 allocs/op
Benchmark/MatrixStringBuilderGrow-8         50000000	        88 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/kevinmartin/go-benchmarks/keyvaluestring	18.208s
```

[CI Results →](https://travis-ci.org/kevinmartin/go-benchmarks)
