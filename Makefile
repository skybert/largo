check: fmt test run
	go mod tidy

run: fmt
	go run .

fmt:
	gofmt -w .

test: fmt
	go test .
