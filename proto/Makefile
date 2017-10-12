all: proto php go
	@echo 'done'

.PHONY: proto
proto:
	protoc --proto_path=./ --php_out=php/build/gen ./example.proto

.PHONY: php
php:
	php php/src/example.php | hexdump -C

.PHONY: go
go:
	@echo 'go build'
