#!/bin/bash

docs_dir="./docs"
[[ ! -d $docs_dir ]] && mkdir $docs_dir

for f in $(find ./api -iname "*.proto"); do
  goctl rpc protoc -m \
    --go_out=. \
    --go-grpc_out=. \
    --zrpc_out=. \
    $f > /dev/null

  protoc --include_imports \
    --proto_path=. \
    --validate_out=paths=source_relative,lang=go:. \
    --openapi_out=paths=source_relative:$docs_dir \
    --descriptor_set_out="${f%.proto}.pb" \
    $f
done
