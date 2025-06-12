default: develop

test:
	go test ./...

develop:
	APP_ENV=dev go run ./cmd/gotodod -debug -config `pwd`/configs

gotodod:
	mkdir -p ./bin
	go build -o ./bin/gotodod ./cmd/gotodod