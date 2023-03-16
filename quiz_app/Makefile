proto:
	rm -f pb/*.dart
	protoc --proto_path=proto --dart_out=pb -Iprotos proto/*.proto google/protobuf/timestamp.proto

.PHONY: proto