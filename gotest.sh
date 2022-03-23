#!/usr/bin/env bash

#配置GOPATH
export GOPATH=~/go
current_dir=$(pwd)

#crypt demo
cd $current_dir/crypt/aes && go test -v .

#ginutil demo
cd $current_dir/ginutil && go test -v .

#net demo
cd $current_dir/net/ding && go test -v .

#async demo
cd $current_dir/task && go test -v .

#types demo
cd $current_dir/types/number && go test -v .
cd $current_dir/types/reflect && go test -v .
cd $current_dir/types/shell && go test -v .
cd $current_dir/types/slice && go test -v .
cd $current_dir/types/times && go test -v .

#uuid demo
cd $current_dir/uuid && go test -v .