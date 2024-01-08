#!/bin/sh

#* variables
PROTO_PATH=./proto
PROTO_OUT=./idl/pb
IDL_PATH=./idl
DOC_OUT=./docs


rm -rf ${PROTO_OUT}
rm transform/*_transformer.go

#! create folders if not eixits
mkdir -p ${DOC_OUT}/html
mkdir -p ${DOC_OUT}/markdown
mkdir -p ${DOC_OUT}/swagger
mkdir -p ${IDL_PATH}
mkdir -p ${PROTO_OUT}

#* gen normal proto
protoc \
	${PROTO_PATH}/*.proto ${PROTO_PATH}/*/*.proto \
	-I=/usr/local/include \
	--proto_path=${PROTO_PATH} \
	--go_out=:${IDL_PATH} \
	--validate_out=lang=go:${IDL_PATH} \
	--go-grpc_out=:${IDL_PATH} \
	--grpc-gateway_out=:${IDL_PATH} \
	--openapiv2_out=:${DOC_OUT}/swagger \
	--custom_out=:${IDL_PATH} \
	--fieldmask_out=lang=go:${PROTO_OUT} \
	--doc_out=:${DOC_OUT}/html --doc_opt=html,index.html

#* gen markdown and tranformer
protoc \
	${PROTO_PATH}/*.proto ${PROTO_PATH}/*/*.proto \
	-I=/usr/local/include \
	--proto_path=${PROTO_PATH} \
	--struct-transformer_out=package=transform,debug=true,goimports=true,helper-package=transformhelpers:. \
	--doc_out=:${DOC_OUT}/markdown --doc_opt=markdown,docs.md

#! remove permission folders
chmod -R 777 ${PROTO_OUT}
chmod -R 777 ${DOC_OUT}