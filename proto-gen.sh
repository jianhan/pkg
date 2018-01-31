#!/bin/bash
echo start to generate protobuf
for file in `find proto | grep -E "\.proto$"`; do
echo generating "$file"
protoc --go_out=plugins=micro,grpc:. $file
done

echo start to inject protobuf
for file in `find proto | grep -E "\.pb.go$"`; do
echo injecting "$file"
protoc-go-inject-tag -input=$file
done
