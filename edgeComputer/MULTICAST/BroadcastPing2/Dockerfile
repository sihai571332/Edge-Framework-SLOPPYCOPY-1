#This will act as container1 grabbing data with bash scripts
FROM golang:1.11-alpine AS builder

#FROM alpine
RUN apk --no-cache add ca-certificates

#install git
#get git command
RUN apk --no-cache add git

#install required components for ping
RUN go get github.com/dmichael/go-multicast/multicast
RUN go install github.com/dmichael/go-multicast/multicast
RUN go get github.com/urfave/cli
RUN go install github.com/urfave/cli

#create directory for ping
RUN mkdir /pinger /go/pinger

#copy pinger's main.go file to pinger directory
COPY /go-multicast-master/examples/pinger/main.go /go/pinger/main.go

#now build pinger's main.go
RUN cd /go/pinger && go build -o main

#install bash
RUN apk --no-cache add bash

ENTRYPOINT ["/go/pinger/main"]