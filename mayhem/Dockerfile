FROM golang:1.21 AS builder

ENV GOTOOLCHAIN=auto

WORKDIR /src
COPY . .

RUN go test -c \
    -gcflags=all=-d=libfuzzer \
    -o /serialization-fuzz.libfuzzer \
    .

FROM ubuntu:22.04
COPY --from=builder /serialization-fuzz.libfuzzer /
