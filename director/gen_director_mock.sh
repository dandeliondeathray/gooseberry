#!/bin/bash

if [ -z "$GOPATH" ]; then
	GOPATH=$HOME/go
fi
$GOPATH/bin/mockgen -destination director_mock/mock.go -package director_mock github.com/dandeliondeathray/gooseberry/director Work