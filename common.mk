BENCHTIME ?= 3s

.PHONY  : bench
bench   :
	@go test -run none -bench . -benchtime $(BENCHTIME) -benchmem

.PHONY  : test
test    :
	@go test -v

.PHONY  : analysis
analysis:
	go build -gcflags '-m -m'

.PHONY  : help
help    :
	@echo "bench    : Runs benchmarks    (\044BENCHTIME=${BENCHTIME})"
	@echo "test     : Runs tests"
	@echo "analysis : Runs optimization analysis"
	@echo "help     : This help"
