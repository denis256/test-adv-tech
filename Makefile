SHELL := /bin/bash
current_dir = $(shell pwd)

all: build


build:
	go mod download
	go build
