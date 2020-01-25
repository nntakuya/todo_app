FROM golang:latest

ENV GO111MODULE=on

RUN mkdir /go/src/work \
    && apt-get update -qq \
    && apt-get install -y mariadb-client vim

WORKDIR /go/src/work

ADD . /go/src/work


