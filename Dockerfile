# We specify the base image we need for our
# go application
FROM golang:1.13.1

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -v
