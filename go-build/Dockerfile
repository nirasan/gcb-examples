# Build Server
FROM golang:latest as builder
ENV GOBIN /go/bin
WORKDIR /go/src/github.com/nirasan/gcb-examples/go-build
COPY / .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

# App Server
FROM alpine:latest
WORKDIR /usr/local/bin
COPY --from=builder /go/src/github.com/nirasan/gcb-examples/go-build/go-build .
ENTRYPOINT ["/usr/local/bin/go-build"]