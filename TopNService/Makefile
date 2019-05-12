all: build docker-build

build:	test
	env GOOS=linux CGO_ENABLED=0 go build -o builds/topnservice
	
test:
	go list ./... | grep -v /vendor/ | xargs go test -v -race -cover

benchmark:
	go list ./... | grep -v /vendor/ | xargs go test -v -run=NOTHING_USE_MAKE_TEST_INSTEAD -benchmem=true -bench=.

run: build
	env LOGGING_LEVEL=debug builds/topnservice


update-deps:
	go mod vendor


docker-build:
	docker build --rm -t topnservice .