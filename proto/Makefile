all: dir proto php go
	@echo 'done'

.PHONY: dir
dir:
	mkdir -p ./php/build/gen ./go/build/gen

.PHONY: proto
proto:
	protoc --proto_path=./ --php_out=php/build/gen ./example.proto
	protoc --proto_path=./ --go_out=go/build/gen ./example.proto

.PHONY: php
php:
	cd ./php && composer install
	php php/src/example.php | hexdump -C

.PHONY: go
go:
	cd ./go && make
	echo 'go build'
