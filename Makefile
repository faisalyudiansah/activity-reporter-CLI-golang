test-cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out && rm -f coverage.out

run:
	go run main.go