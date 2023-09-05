PKGS := ./...

.PHONY: check
check:
	go vet $(PKGS)
	go test $(PKGS)

.PHONY: tidy
tidy:
	go clean $(PKGS)
	go fmt $(PKGS)
	go mod tidy
