dev:
	CompileDaemon -build="go build -o ./dist/go-restapi" -command="./dist/go-restapi"

build:
	go build -o ./dist/go-restapi

run:
	make build && ./dist/go-restapi