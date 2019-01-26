.DEFAULT_GOAL := start

start: build execute

build:
	go build

execute:
	./star-wars-go
