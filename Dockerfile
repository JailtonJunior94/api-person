FROM golang:1.15.2-alpine

WORKDIR /go/src/

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build main.go

EXPOSE 3000
ENTRYPOINT ["./main"]