FROM golang:1.13.1-alpine3.10 AS builder
WORKDIR /go/src/authreq
COPY ./ ./
RUN GOOS=linux go build -o protected ./cmd/protected
RUN GOOS=linux go build -o caddy ./cmd/caddy

FROM alpine:3.10 AS protected
WORKDIR /root/
COPY --from=builder /go/src/authreq/protected .

FROM alpine:3.10 AS caddy
WORKDIR /root/
COPY --from=builder /go/src/authreq/caddy .
COPY --from=builder /go/src/authreq/cmd/caddy/config.json .
