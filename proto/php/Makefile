all:
	protoc --proto_path=.. --php_out=build/gen ../example.proto
	php src/example.php | hexdump -C
