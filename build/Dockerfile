FROM golang:1.13.1-alpine3.10 AS builder
WORKDIR /go/src/oauth-demo
COPY ./ ./
RUN GOOS=linux go build -o protected ./cmd/protected
RUN GOOS=linux go build -o caddy ./cmd/caddy

FROM alpine:3.10 AS protected
WORKDIR /root/
COPY --from=builder /go/src/oauth-demo/protected .

FROM alpine:3.10 AS caddy
WORKDIR /root/
COPY --from=builder /go/src/oauth-demo/caddy .
COPY --from=builder /go/src/oauth-demo/cmd/caddy/config.json .
