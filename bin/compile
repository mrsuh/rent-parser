#!/bin/sh

cd "$( cd `dirname $0` && pwd )/.."

go get -u github.com/valyala/fasthttp
go get gopkg.in/yaml.v2
go get github.com/mrsuh/cli-config

go build -o bin/main -i src/main.go
