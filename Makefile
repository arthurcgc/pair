.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: lint
lint:
	@echo "Running linter..."
	@golangci-lint run

.PHONY: tag
tag:
	@echo "Creating tag..."
	if [ -z "$(VERSION)" ]; then echo "VERSION is not set"; exit 1; fi
	@git tag -a $(VERSION) -m "Release $(VERSION)"
	@git push origin $(VERSION)