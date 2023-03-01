export GO111MODULE=on
export GOPROXY=https://goproxy.io

all:
	swag init
	mkdir -p bin
	go build -o ./bin/main main.go 
default: all

clean:
	rm -f ./bin/main
