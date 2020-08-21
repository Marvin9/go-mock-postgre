.PHONY: test

test:
	DATABASE_NAME=test_original go test ./... -p 1