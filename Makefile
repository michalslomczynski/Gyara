all: build install

build:
	go build -o gyara cli/main.go

install:
	sudo mv gyara /usr/local/bin/

clean:
	rm gyara | sudo rm /usr/local/bin/gyara

.PHONY: all build install clean