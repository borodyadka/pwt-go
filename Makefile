.PHONY: test

compile_proto = protoc \
	-I/usr/local/include \
	-I$(shell go env GOPATH)/src \
	-I. \
	--gogoslick_out=plugins=grpc,paths=source_relative,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:.\
	./$(1).proto

token.pb.go: token.proto
	$(call compile_proto,token)
proto: token.pb.go

test: proto
	go test -mod vendor ./...
