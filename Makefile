.DEFAULT_GOAL := all
BIN_NAME = "exirc"

all: build compress

build:
	@go build -o $(BIN_NAME) -ldflags="-s -w" *.go

compress: 
	@upx $(BIN_NAME)
