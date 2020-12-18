SRC_PATH=$(GOPATH)/src/github.com/nfqphuong/ddos

build:
	go build -o bin/ddos $(SRC_PATH)/main.go