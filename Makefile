GO=$(shell which go)
WRK=$(shell which wrk)
test:
	$(GO) test
benchmark:
	$(GO) test -test.bench=".*"
fe-dev:
	http-server fe -p 8080
http-bm-fe:
	$(WRK) -t10 -c100 -d20s http://localhost:8080/
compile:
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o tinyurl
clean:
	rm tinyurl
