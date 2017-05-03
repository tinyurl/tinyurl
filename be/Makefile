GO=$(shell which go)
GB=$(shell unalias gb;which gb)

build:
	@echo $(GB)
	$(GB) build tinyurl

dev: build
	nohup ./bin/tinyurl -dbname testurl -user root -pass rootLu -dbport 2306 > tinyurl.log 2>&1 &
test: build
	./bin/tinyurl -dbname testurl -user root -pass rootLu -dbport 2306
