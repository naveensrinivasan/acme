#!/usr/bin/env bash
go build \
    -o acme \
    -a \
    -installsuffix \
    cgo \
    -tags netgo \
    "./cmd"
