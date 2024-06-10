all: build install

build:
	go build -o gyara cli/main.go

install:
	sudo mv gyara /usr/local/bin/