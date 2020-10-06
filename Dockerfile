FROM golang:1.14.4-alpine

WORKDIR /go/src/

COPY . .

RUN GOOS=linux go build main.go

EXPOSE 3000
ENTRYPOINT ["./main"]