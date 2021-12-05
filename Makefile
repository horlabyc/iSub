run:
	nodemon --exec go run main.go --signal SIGTERM
build:
	go build -o bin/main main.go