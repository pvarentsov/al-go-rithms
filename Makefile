.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@echo
	@echo "   stupid-sort   - Run stupid sorting demo"
	@echo "   bubble-sort   - Run bubble sorting demo"
	@echo "   shaker-sort   - Run shaker sorting demo"
	@echo
	@echo "   fmt           - Format source code"

.PHONY: stupid-sort
stupid-sort:
	go run ./cmd/stupid_sorting_demo.go

.PHONY: bubble-sort
bubble-sort:
	go run ./cmd/bubble_sorting_demo.go
	
.PHONY: shaker-sort
shaker-sort:
	go run ./cmd/shaker_sorting_demo.go

.PHONY: comb-sort
comb-sort:
	go run ./cmd/comb_sorting_demo.go

.PHONY: quick-sort
quick-sort:
	go run ./cmd/quick_sorting_demo.go

.PHONY: fmt
fmt:
	go fmt ./...

.DEFAULT_GOAL := help
