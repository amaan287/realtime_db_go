build:
	@go build -o bin/real_time_db cmd/main.go


run: build
	@./bin/real_time_db
test: 
	go test -v ./..