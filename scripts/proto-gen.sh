#!/bin/bash
CURRENT_DIR=$(pwd)
for i in $(find ${CURRENT_DIR}/protos/* -type d); do
#  echo $i
  protoc -I=${i} -I=${CURRENT_DIR}/protos -I /usr/local/include --proto_path=${CURRENT_DIR}/protos --go_out=plugins=grpc:${CURRENT_DIR} ${i}/*.proto
done