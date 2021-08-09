SHELL := /bin/bash
current_dir = $(shell pwd)

all: build


build:
	go mod download
	go build

container:
	docker build . -t denis256/test-adv-tech:$(shell git rev-parse --short HEAD)
	docker push denis256/test-adv-tech:$(shell git rev-parse --short HEAD)
