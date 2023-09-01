FROM docker.io/golang:1.21 AS builder

WORKDIR /src
COPY go.mod go.sum /src
RUN go mod download
COPY . /src
RUN go build

FROM docker.io/debian:bookworm

COPY --from=builder /src/example /usr/local/bin/example

CMD ["/usr/local/bin/example"]
