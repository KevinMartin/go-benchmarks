# Fingerprint Validator

This benchmark tests different ways to validate a fingerprint of arbitrary length.

## Expectations

Given a string, validate the string contains hex pairs separated by colons. Return true if valid, false otherwise.

```
given = {
  ex1 = "12:23:34:45"
  ex2 = "foo:bar:zzz"
}

expect = {
  ex1 = true
  ex2 = false
}
```

## Example Use-Cases
- Validate Public Key fingerprints as input.

## Running the Benchmark

```
$ make bench
```

## Results

```
goos: darwin
goarch: amd64
pkg: github.com/kevinmartin/go-benchmarks/fingerprint-validator
Benchmark/ValidateStringsSplitStringsCount_small-8         	30000000	       151 ns/op	      64 B/op	       1 allocs/op
Benchmark/ValidateStringsSplitStringsCount_medium-8        	20000000	       245 ns/op	     128 B/op	       1 allocs/op
Benchmark/ValidateStringsSplitStringsCount_large-8         	10000000	       444 ns/op	     256 B/op	       1 allocs/op
Benchmark/ValidateStringsSplitStrings_small-8              	30000000	       145 ns/op	      64 B/op	       1 allocs/op
Benchmark/ValidateStringsSplitStrings_medium-8             	20000000	       237 ns/op	     128 B/op	       1 allocs/op
Benchmark/ValidateStringsSplitStrings_large-8              	10000000	       459 ns/op	     256 B/op	       1 allocs/op
Benchmark/ValidateLoopPairsCount_small-8                   	100000000	       38.8 ns/op	       0 B/op	       0 allocs/op
Benchmark/ValidateLoopPairsCount_medium-8                  	50000000	       73.5 ns/op	       0 B/op	       0 allocs/op
Benchmark/ValidateLoopPairsCount_large-8                   	30000000	       149 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/kevinmartin/go-benchmarks/fingerprint-validator	41.640s
```

[CI Results →](https://travis-ci.org/kevinmartin/go-benchmarks)
