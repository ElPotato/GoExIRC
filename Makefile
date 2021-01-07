.DEFAULT_GOAL := all
BIN_NAME := "exirc"
LD_FLAGS := -ldflags='-X "main.ID=YXZZZZ3" -s -w'

all: build compress

build:
	@go build $(LD_FLAGS) -o $(BIN_NAME) *.go

compress: 
	@upx $(BIN_NAME)
