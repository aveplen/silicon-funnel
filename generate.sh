#!/bin/sh

rm -rf pkg
mkdir -p pkg

protoc -I api \
	--go_out pkg \
	--go_opt paths=source_relative \
	--go-grpc_out pkg \
	--go-grpc_opt paths=source_relative \
	--grpc-gateway_out pkg \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	--grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out pkg \
	--openapiv2_opt logtostderr=true \
	api/imap_concentrator/v1/imap_concentrator.proto

P=$(pwd)
cd pkg/imap_concentrator/v1
minimock
cd $P
