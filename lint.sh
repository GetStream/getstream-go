#!/bin/bash

go fmt .
go vet .
go install mvdan.cc/gofumpt@latest
gofumpt -l -w .