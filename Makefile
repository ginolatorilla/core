TEST_REGEX=".*"
TEST_PACKAGE="./..."

.PHONY: all
all: test

.PHONY: test
test: tidy
	@echo "🌡  Running tests..."
	go test -race $(BUILD_FLAGS) $(LD_FLAGS) -run $(TEST_REGEX) $(TEST_PACKAGE)

.PHONY: test/cover
test/cover: tidy
	@echo "🌡️  Running tests..."
	@go test -coverprofile=/tmp/coverage.out -race $(BUILD_FLAGS) $(LD_FLAGS) -run $(TEST_REGEX) $(TEST_PACKAGE)
	@go tool cover -html=/tmp/coverage.out

.PHONY: tidy
tidy:
	@echo "🧹 Tidying up package dependencies..."
	go mod tidy

.PHONY: doc
doc:
	@go install golang.org/x/pkgsite/cmd/pkgsite@latest
	@pkgsite -open

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  help       - Show this help message"
	@echo "  all        - Run test and tidy (default)"
	@echo "  test       - Run tests"
	@echo "  test/cover - Run tests with coverage"
	@echo "  tidy       - Sort out package dependencies"
	@echo "  doc        - Open the documentation in the browser"
