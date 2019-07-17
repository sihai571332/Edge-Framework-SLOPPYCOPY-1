#use golang as a base image
FROM golang:1.11-alpine AS builder


#install bash
RUN apk --no-cache add bash

#create a pinger directory
RUN mkdir /pinger /go/pinger

#create listener directory
RUN mkdir /listener /go/listener

#copy pinter/main.go to pinger directory
COPY /go-multicast-master/examples/pinger/main.go /go/pinger/main.go

#copy listener/main.go to listener directory
COPY /go-multicast-master/examples/listener/main.go /go/listener/main.go

#install git command
RUN apk --no-cache add git

#get and install packages
RUN go get github.com/dmichael/go-multicast/multicast && go install github.com/dmichael/go-multicast/multicast
RUN go get github.com/urfave/cli && go install github.com/urfave/cli

#build pinger's main.go, listener's main.go, and startup's hello.go
RUN cd /go/pinger && go build -o main
RUN cd /go/listener && go build -o main


ADD . /src

RUN cd /src \
 && go build -o hello

#take care of cache
FROM alpine
RUN apk --no-cache add ca-certificates

COPY --from=builder /src/hello /bin


EXPOSE 9000
ENTRYPOINT ["hello"] 