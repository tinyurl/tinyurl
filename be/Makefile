GO=$(shell which go)
GB=$(shell unalias gb;which gb)\

build:
	$(GB) build tinyurl

test: build
	./bin/tinyurl -dbname testurl -user root -pass root -dbport 3306
