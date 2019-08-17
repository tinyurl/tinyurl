GO_BUILD_FLAGS=
PKGS=$(shell go list ./... | grep -E -v "(vendor)")

all:
	go build $(GO_BUILD_FLAGS) -o tinyurl

build:
	go build $(GO_BUILD_FLAGS) -o tinyurl
build-image:
	docker build -t adolphlwq/tinyurl .

# run binary in dev mode
dev:
	./tinyurl -config dev.properties

# run test cases
test-mysql:
	TINYURL_CONFIG_PATH="./store/test_data/mysql.properties" go test --cover $(PKGS)
test-sqlite3:
	TINYURL_CONFIG_PATH="./store/test_data/sqlite3.properties" go test --cover $(PKGS)

test:
	echo "test with mysql"
	make test-mysql
	echo "test with sqlite3"
	make test-sqlite3

start-container:
	docker run -d --name tinyurl_mysql --net host \
	-e MYSQL_ALLOW_EMPTY_PASSWORD=yes -e MYSQL_DATABASE=tinyurldb mysql:5.7

clean-container:
	docker stop tinyurl_mysql
	docker rm tinyurl_mysql

# run benchmark test
benchmark:
	go test -test.bench=".*"

# run frontend with http-server
fe-dev:
	http-server frontend -p 8080
http-bm-fe:
	wrk -t10 -c100 -d20s http://localhost:8080/

clean:
	rm -f tinyurl