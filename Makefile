GO=$(shell which go)

test:
	$(GO) test
benchmark:
	$(GO) test -test.bench=".*"
fe-dev:
	http-server fe
compile:
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o tinyurl
clean:
	rm tinyurl
