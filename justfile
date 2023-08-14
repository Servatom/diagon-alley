#!/usr/bin/env just --justfile

set dotenv-load

GO := "go"

GOVET_COMMAND := GO + " vet"
GOTEST_COMMAND := GO + " test"
GOCOVER_COMMAND := GO + " tool cover"
GOBUILD_COMMAND := GO + " build"
GORUN_COMMAND := GO + " run"

# display all commands
default:
    @just --list --unsorted

# Run static checks
check:
    {{GOVET_COMMAND}} ./...

# Clean dist directory and rebuild the binary file
build:
    rm -rf ./dist
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 {{GOBUILD_COMMAND}} -ldflags="-w -s" -o ./dist/diagon-alley ./src
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 {{GOBUILD_COMMAND}} -ldflags="-w -s" -o ./dist/diagon-alley-create-user ./src/delivery/cmd/auth/main.go
# create super user
create-user:
    {{GORUN_COMMAND}} ./src/cmd/auth/main.go --create-user
