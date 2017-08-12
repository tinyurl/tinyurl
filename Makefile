GO_BUILD_FLAGS=
PKGS=$(shell go list ./... | grep -E -v "(vendor)")

all:
	go build $(GO_BUILD_FLAGS) -o tinyurl
dev:
	./tinyurl -config dev.properties
test:
	go test --cover $(PKGS)
benchmark:
	go test -test.bench=".*"
fe-dev:
	http-server frontend -p 8080
http-bm-fe:
	wrk -t10 -c100 -d20s http://localhost:8080/
clean:
	rm -f tinyurl