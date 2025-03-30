.PHONY: test test100

test:
	@go test -cover -count 1 -failfast ./

test100:
	@go test -cover -count 100 -failfast ./
