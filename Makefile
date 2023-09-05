PKGS := ./...

.PHONY: check
check:
	go vet $(PKGS)
	go test $(PKGS)

.PHONY: clean
clean:
	go clean $(PKGS)
