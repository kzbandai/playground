all: dir proto php go
	@echo 'done'

.PHONY: dir
dir:
	mkdir -p \
	./php/build/pb ./go/build/pb ./java/build/pb ./python/build/pb ./ruby/build/pb

.PHONY: proto
proto:
	protoc --proto_path=./ --php_out=php/build/pb ./example.proto
	protoc --proto_path=./ --go_out=go/build/pb ./example.proto
	protoc --proto_path=./ --java_out=java/build/pb ./example.proto
	protoc --proto_path=./ --ruby_out=ruby/build/pb ./example.proto
	protoc --proto_path=./ --python_out=python/build/pb ./example.proto

.PHONY: php
php:
	cd ./php && composer install
	php php/src/example.php | hexdump -C

.PHONY: go
go:
	cd ./go && make
	./go/main | hexdump -C
