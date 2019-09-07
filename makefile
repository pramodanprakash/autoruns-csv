#!/usr/bin/make -f

compile:
	go build ./...

build: compile

.PHONY: compile build