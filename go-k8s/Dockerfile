FROM alpine:latest
WORKDIR /usr/local/bin/

COPY gopath/bin/go-k8s .

ENTRYPOINT ["/usr/local/bin/go-k8s"]