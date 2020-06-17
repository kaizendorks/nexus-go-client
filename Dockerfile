FROM golang:1.14.0-alpine

RUN apk update \
    && apk add \
      build-base \
      graphviz \
      ttf-freefont

WORKDIR /go/src/github.com/kaizendorks/nexus-go-client
COPY go.mod go.sum ./
RUN go mod download
