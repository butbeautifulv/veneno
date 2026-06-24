.PHONY: help test-engage test-engage-unit

help:
	@echo "Targets:"
	@echo "  test-engage       Run engage unit tests"
	@echo "  test-engage-unit  Alias for test-engage"

test-engage test-engage-unit:
	@cd engage/serve && go test ./...
