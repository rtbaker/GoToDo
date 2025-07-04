default: develop

test:
	go test `go list ./... | grep -v GoToDo/Model`

coverage:
	go test -cover `go list ./... | grep -v GoToDo/Model`

start-develop:
	docker compose -f development/docker-compose-dev.yml up -d

stop-develop:
	docker compose -f development/docker-compose-dev.yml down
	
develop:
	APP_ENV=dev go run ./cmd/gotodod -debug -config `pwd`/configs

develop-inmemory:
	APP_ENV=dev-inmemory go run ./cmd/gotodod -debug -config `pwd`/configs

gotodod:
	mkdir -p ./bin
	go build -o ./bin/gotodod ./cmd/gotodod

hashpassword:
	mkdir -p ./bin
	go build -o ./bin/hashpassword ./cmd/hashpassword