
build:
	go build -o main

run:
	./main

clean:
	go clean

lint:
	go fmt .

all: clean lint build run
