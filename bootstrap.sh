#!/usr/bin/env bash
rm -rf ./vendor
go get -u github.com/golang/dep/cmd/dep
dep ensure