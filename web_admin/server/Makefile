export GO111MODULE=on

all:
	swag init
	mkdir -p bin
	go build -o ./bin/main main.go 
default: all

clean:
	rm -f ./bin/main
