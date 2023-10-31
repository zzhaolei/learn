#!/bin/bash

protos=$(find ./api -iname "*.proto")
if [[ -z $protos ]]; then
  echo "Not found proto file." && exit
fi

# Generate go-zero
for f in $protos; do
  goctl rpc protoc -m \
    --go_out=. \
    --go-grpc_out=. \
    --zrpc_out=. \
    $f > /dev/null
done

# Generate docs/validate/descriptor
docs_dir="./docs"
[[ ! -d $docs_dir ]] && mkdir $docs_dir
protoc --include_imports \
  --proto_path=. \
  --validate_out=paths=source_relative,lang=go:. \
  --openapi_out=paths=source_relative:$docs_dir \
  --descriptor_set_out="./api/http.pb" \
  $protos


