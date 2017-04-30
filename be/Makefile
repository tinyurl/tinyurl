GO=$(shell which go)
GB=$(shell unalias gb;which gb)

build:
	@echo $(GB)
	$(GB) build tinyurl

test: build
	./bin/tinyurl -dbname testurl -user root -pass root -dbport 3306
