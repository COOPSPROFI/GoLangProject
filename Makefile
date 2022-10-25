run:
	go run cmd/api/main.go

build:
	go build -o ../build/main cmd/api/main.go

daemon:
	CompileDaemon -build="go build -o ./build/main cmd/api/main.go" -command="./build/main"