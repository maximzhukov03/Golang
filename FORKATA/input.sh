#!/bin/bash

if ["$#" -ne 1]; then echo "Module name argument is missing"
    exit 1
fi

go mod init $1
go get github.com/yuin/goldmark
go mod tidy