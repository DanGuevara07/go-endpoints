.PHONY: build clean
build: clean
	env GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/build ./cmd/server/main.go
clean:
	rm -rf ./bin

start: build
	node ./node_modules/.bin/serverless offline start --useDocker --dockerNetwork go-endpoints_default --host 0.0.0.0

test: 
	@echo "no tests yet"