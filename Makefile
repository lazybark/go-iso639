.PHONY: test

test:
	go run gotest.tools/gotestsum@latest --format-icons hivis --format-hide-empty-pkg

lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.64 golangci-lint run
