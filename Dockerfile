FROM golang:1.18-alpine AS builder
WORKDIR /app/article-api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o article-api article-api/cmd

FROM alpine
RUN apk update && apk add --no-cache tzdata
WORKDIR /app
COPY --from=builder /app/article-api /app/
ENTRYPOINT ["./article-api"]