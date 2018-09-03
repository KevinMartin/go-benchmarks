BENCHTIME ?= 3s

.PHONY: bench
bench :
	@go test -run none -bench . -benchtime $(BENCHTIME) -benchmem

.PHONY: test
test  :
	@go test -v

.PHONY: help
help  :
	@echo "bench : Runs benchmarks    (\044BENCHTIME=${BENCHTIME})"
	@echo "test  : Runs tests"
	@echo "help  : This help"
