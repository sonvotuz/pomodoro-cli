build:
	@go build -o out

test:
	@go test -v ./...
	
run: build
	@./out