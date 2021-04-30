FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app

COPY . .

RUN go clean --modcache
RUN GOOS=linux go build -o main main.go

FROM alpine:3.13
WORKDIR /app

RUN apk --no-cache add tzdata

ENV TZ=America/Sao_Paulo
ENV PORT=3000

COPY --from=builder /app/main .
COPY .env .

EXPOSE 3000

CMD [ "/app/main" ]