FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8081

ENTRYPOINT [ "./app" ]