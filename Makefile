.PHONY: run
run: build
	./ddb-backend-developer-challenge

.PHONY: test
test: install_dev
	go test ./...
	go mod tidy

.PHONY: build
build: install
	go build

.PHONY: install
install:
	go get -t ./...
	go mod tidy

.PHONY: install_dev
install_dev: install
	go get -tags tools ./...

.PHONY: gen
gen: install_dev
	go generate ./...






