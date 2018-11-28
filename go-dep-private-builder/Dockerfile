FROM alpine:latest
WORKDIR /usr/local/bin/

COPY gopath/bin/go-dep-private-builder .

ENTRYPOINT ["/usr/local/bin/go-dep-private-builder"]