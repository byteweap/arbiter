.PHONY: lint test cover coverage-html release

lint:
	golangci-lint run ./...

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -func=coverage.out

coverage-html:
	go test -coverprofile=coverage.out -covermode=atomic ./...
	go tool cover -html=coverage.out -o coverage.html

release:
	@VERSION="$(word 2,$(MAKECMDGOALS))"; \
	if [ -z "$$VERSION" ]; then echo "Usage: make release v1.0.1"; exit 1; fi; \
	if git rev-parse "$$VERSION" >/dev/null 2>&1; then echo "Error: Tag $$VERSION already exists"; exit 1; fi; \
	git tag -a $$VERSION -m "Release $$VERSION"; \
	git push origin master --tags; \
	echo "Released $$VERSION successfully!"

%:
	@:
