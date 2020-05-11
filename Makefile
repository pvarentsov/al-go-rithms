.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@echo
	@echo "   stupid-sort   - Run stupid sorting demo"
	@echo "   fmt           - Format source code"

.PHONY: stupid-sort
stupid-sort:
	go run ./cmd/stupid_sorting_demo.go

.PHONY: bubble-sort
bubble-sort:
	go run ./cmd/bubble_sorting_demo.go

.PHONY: fmt
fmt:
	go fmt ./...

