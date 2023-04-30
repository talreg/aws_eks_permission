default:
	go build -o bin/runner
run: default
	@echo
	bin/runner
clean:
	@rm -rf bin
test:
	@go clean -testcache
	@go test ./...
dockerize: default
	docker build -t secrets-tester .