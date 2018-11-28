FROM alpine:latest
WORKDIR /usr/local/bin/

COPY gopath/bin/go-dep-builder .

ENTRYPOINT ["/usr/local/bin/go-dep-builder"]