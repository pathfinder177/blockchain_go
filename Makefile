.PHONY: test test100

testDir=./tests

test:
	@go test -cover -count 1 -failfast "${testDir}"

test100:
	@go test -cover -count 100 -failfast "${testDir}"
