FROM golang:latest

RUN GO111MODULE=on go get github.com/golang/mock/mockgen@v1.4.4

WORKDIR /src

ENTRYPOINT [ "go", "generate", "-v", "./..." ]