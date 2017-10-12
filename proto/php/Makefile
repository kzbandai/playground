all:
	protoc --proto_path=src --php_out=build/gen src/example.proto
	php test.php | hexdump -C
