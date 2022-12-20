test:
	go test -v ./...
	rm -rf test_template

test-coverage:
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile coverage.out
	go tool cover -html coverage.out -o coverage.html
	rm -rf test_template