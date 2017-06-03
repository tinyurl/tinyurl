GO=$(shell which go)

test:
	echo "test"
	echo $(GO)

benchmark:
	$(GO) test -test.bench=".*"
