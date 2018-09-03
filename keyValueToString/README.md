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
pkg: github.com/KevinMartin/go-benchmarks/keyValueToString
Benchmark/MapWithSort-4         	10000000	       480 ns/op	      96 B/op	       5 allocs/op
Benchmark/Matrix-4              	20000000	       237 ns/op	      64 B/op	       4 allocs/op
PASS
ok  	github.com/KevinMartin/go-benchmarks/keyValueToString	10.326s
```
