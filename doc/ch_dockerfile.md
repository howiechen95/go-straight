# Dockerfile

- [Gin实践 连载九 将Golang应用部署到Docker](https://segmentfault.com/a/1190000013960558)

文件命名必须为： Dockerfile

```
FROM golang:alpine AS development
WORKDIR $GOPATH/src/test
COPY . .
RUN go build -o test

FROM alpine:latest AS production
WORKDIR /root/test/
COPY --from=development /go/src/test .
EXPOSE 8080
ENTRYPOINT ["./test"]
```

```
FROM golang:latest

WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
```