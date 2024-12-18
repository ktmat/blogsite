BIN=Blogsite
SRC=$(wildcard *.go)

all: build

build: $(SRC)
	@echo "Pulling latest from GitHub..."
	@git pull
	@echo "Building $(BIN)..."
	@go build -o $(BIN) main.go

run:
	@echo "Running $(BIN) Application!"
	@./$(BIN)

clean:
	@echo "Cleaning up..."
	@rm -f $(BIN)

runclean: clean build run
	@echo "Cleaned, built, and running $(BIN)"
