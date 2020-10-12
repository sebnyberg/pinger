FROM golang:1.15-alpine3.12 AS base

WORKDIR /code

COPY . .

RUN go mod download

RUN go build -o main .

FROM alpine:3.12 as release

COPY --from=base /code/main /go/bin/main

ENTRYPOINT ["/go/bin/main"]