.PHONY: check
check:
	go vet
	go test ./...
