# Build Server
FROM golang:latest as builder
ENV GOBIN /go/bin
WORKDIR /go/src/github.com/nirasan/gcb-examples
COPY / .
WORKDIR /go/src/github.com/nirasan/gcb-examples/go-dep-build
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure
RUN go test
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

# App Server
FROM alpine:latest
WORKDIR /usr/local/bin
COPY --from=builder /go/src/github.com/nirasan/gcb-examples/go-dep-build/go-dep-build .
ENTRYPOINT ["/usr/local/bin/go-dep-build"]