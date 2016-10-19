default: all

all: go docker

go: main.go
	env GOOS=linux GOARCH=amd64 go build

docker: Dockerfile
	docker build -t jekyll-build .

