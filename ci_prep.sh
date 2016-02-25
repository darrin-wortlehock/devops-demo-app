#!/bin/bash
go get github.com/stretchr/testify/assert
eval "$(chef shell-init bash)"
bundle install
