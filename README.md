[![Go Report Card](https://goreportcard.com/badge/github.com/kevinmartin/go-benchmarks)](https://goreportcard.com/report/github.com/kevinmartin/go-benchmarks)
[![Build Status](https://travis-ci.org/kevinmartin/go-benchmarks.svg?branch=master)](https://travis-ci.org/kevinmartin/go-benchmarks)

# go-benchmarks

A collection of random Go benchmarks.

TLDR: [View Latest Benchmark Results on CI →](https://travis-ci.org/kevinmartin/go-benchmarks)

## How to Use

Each folder contains a benchmark I've done in the past with code, tests, benchmarks, and results.

Feel free to clone and run the benchmarks on your own. There is a standard way of running each benchmark.

## Commands in each folder

Each folder has a `Makefile` with the following commands available:

### Running Benchmarks

```
$ make bench
```

### Running Tests

```
$ make test
```

### Running Escape Analysis

```
$ make analysis
```

### Help?

```
$ make help
```
