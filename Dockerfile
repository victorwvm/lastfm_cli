FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .
RUN go mod tidy

RUN go build -o lastfm_cli main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/lastfm_cli .


ENTRYPOINT ["./lastfm_cli"]