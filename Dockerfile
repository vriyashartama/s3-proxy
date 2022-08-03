ARG GOLANG_VERSION=1.17.5
FROM golang:${GOLANG_VERSION}-buster as builder

WORKDIR ${GOPATH}/src/github.com/liehart/s3-proxy

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test -v

RUN go build -a \
    -o ${GOPATH}/bin/s3-proxy

FROM debian:buster-slim

COPY --from=builder /go/bin/s3-proxy /usr/local/bin/s3-proxy

ENV PORT 9990

#USER nobody

ENTRYPOINT ["/usr/local/bin/s3-proxy"]

EXPOSE ${PORT}