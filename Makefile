GO=$(shell which go)

test:
	echo "test"
	echo $(GO)
fe-dev:
	http-server fe
compile:
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o tinyurl
benchmark:
	$(GO) test -test.bench=".*"
clean:
	rm tinyurl
