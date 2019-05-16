# Golang image for building the binary
FROM golang:alpine AS builder
LABEL maintainer "John McKenzie <jmckind@gmail.com>"

WORKDIR /go/src/github.com/jmckind/gtrchive
COPY *.go .
RUN set -x && \ 
    apk add git --no-cache && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gtrchive .

# Minimal image for final binary
FROM scratch
LABEL maintainer "John McKenzie <jmckind@gmail.com>"

COPY --from=builder /go/src/github.com/jmckind/gtrchive/gtrchive /gtrchive

WORKDIR /
CMD ["/gtrchive"]
