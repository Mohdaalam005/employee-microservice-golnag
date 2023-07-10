# syntax=docker/dockerfile:1

FROM golang:1.19

RUN mkdir /App

ADD . /App

WORKDIR /App

RUN go build -o main .


CMD ["/App/main"]