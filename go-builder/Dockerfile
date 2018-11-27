FROM alpine:latest
WORKDIR /usr/local/bin/
COPY gopath/bin/go-builder .

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/go-builder"]