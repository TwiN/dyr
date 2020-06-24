#!/bin/bash

APPLICATION_NAME="dyr"

mkdir bin
rm -rf bin/${APPLICATION_NAME}-*.zip

env GOOS=darwin GOARCH=amd64 go build -mod vendor -o ${APPLICATION_NAME}
chmod +x ${APPLICATION_NAME}
zip bin/${APPLICATION_NAME}-darwin64.zip ${APPLICATION_NAME} -m

env GOOS=linux GOARCH=amd64 go build -mod vendor -o ${APPLICATION_NAME}
chmod +x ${APPLICATION_NAME}
zip bin/${APPLICATION_NAME}-linux64.zip ${APPLICATION_NAME} -m

env GOOS=windows GOARCH=amd64 go build -mod vendor -o ${APPLICATION_NAME}.exe
chmod +x ${APPLICATION_NAME}.exe
zip bin/${APPLICATION_NAME}-windows64.zip ${APPLICATION_NAME}.exe -m
