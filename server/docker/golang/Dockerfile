FROM golang:1.11-alpine

ARG go_release=release

ENV CGO_ENABLED 0
ENV GIN_MODE ${go_release}

RUN apk add --upgrade --no-cache git make \
  && go get github.com/cespare/reflex \
  && mkdir /go/app

WORKDIR /go/app
COPY . /go/app

RUN make build

CMD ["./bin/todone-server"]
