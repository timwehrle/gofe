.PHONY: test
test:
	go test -v ./...

.PHONY: test/cover
test/cover:
	go test -v -coverprofile=c.out ./...
	go tool cover -html=c.out

.PHONY: tidy
tidy:
	go mod tidy -v
	go fmt ./...

.PHONY: audit
audit: test
	go mod tidy
	go mod verify
	go fmt ./...
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: build
build:
	goreleaser --snapshot --clean