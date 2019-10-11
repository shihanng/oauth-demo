FROM golang:1.13.1-alpine3.10 AS builder
WORKDIR /go/src/oauth-demo
COPY ./ ./
RUN GOOS=linux go build -o protected ./cmd/protected

FROM alpine:3.10
WORKDIR /root/
COPY --from=builder /go/src/oauth-demo/protected .
CMD ["./protected"]
