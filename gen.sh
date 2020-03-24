#!/usr/bin/env bash

PROTO_SRC_PATH=./
IMPORT_MAPPING="commo.proto=protoc-gen-markdown/common"

#find . -name '*.proto' -exec protoc --proto_path=$PROTO_SRC_PATH \
#--twirp_out=prefix=vocaldh,M$IMPORT_MAPPING:$PROTO_SRC_PATH \
#--go_out=M$IMPORT_MAPPING:$PROTO_SRC_PATH  {} \;


go build -o protoc-gen-markdown ./

# find ./ -name '*.proto' -exec \
# protoc \
# --plugin=protoc-gen-markdown=/Users/MS/Documents/goworkspace/src/protoc-gen-markdown/protoc-gen-markdown \
# --markdown_out=path_prefix=/vocaldh,M$IMPORT_MAPPING:. {} \;

protoc --plugin=protoc-gen-markdown=/Users/MS/Documents/goworkspace/src/protoc-gen-markdown/protoc-gen-markdown --markdown_out=path_prefix=/vocaldh,M$IMPORT_MAPPING:. hello.proto
