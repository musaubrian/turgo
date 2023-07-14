GO:=go
BIN:= turgo
SRC:= .
BIN_LOC:=./bin/$(BIN)
BUILD:=$(GO) build -o $(BIN_LOC) $(SRC)

build:
	@$(BUILD)
run:build
	@$(BIN_LOC)

topics:build
	@$(BIN_LOC) topics

list:build
	@$(BIN_LOC) list
