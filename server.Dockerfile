# syntax=docker/dockerfile:1

FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd/server
RUN go build -o main .
RUN git clone https://github.com/vishnubob/wait-for-it.git
