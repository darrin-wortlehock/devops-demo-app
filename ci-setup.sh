#!/bin/bash
set -e
export GOPATH=$HOME/golang
mkdir -p $GOPATH
go get github.com/stretchr/testify/assert
eval "$(chef shell-init bash)"
bundle install
