#!/usr/bin/env bash

set -e

PROJECT_DIR=$(cd "$(dirname "$0")/.." && pwd)
GENERATED_DIR="${PROJECT_DIR}/pkg/bank/protobuf"

PROTO_PATHS=(
  "${PROJECT_DIR}/api/protos"
)

rm -rf "${GENERATED_DIR}"
mkdir -p "${GENERATED_DIR}"

arg_proto_files=""
arg_include_paths=""
for proto_path in "${PROTO_PATHS[@]}"; do
  arg_include_paths="${arg_include_paths} -I=${proto_path}"
  arg_proto_files="${arg_proto_files} $(find "${proto_path}" -iname "*.proto" -print0 | xargs -0)"
done
eval protoc ${arg_include_paths} \
  --go_out=paths=source_relative:"${GENERATED_DIR}" \
  --go-grpc_out=paths=source_relative:"${GENERATED_DIR}" \
  ${arg_proto_files}
