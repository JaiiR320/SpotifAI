all: templ
	@go build -o bin/app.exe cmd/main.go
	@./bin/app.exe

templ:
	@templ generate

test:
	@go test -v ./...

tailwind:
	@tailwind -i view/input.css -o static/output.css --watch
	