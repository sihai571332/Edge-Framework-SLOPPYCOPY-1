#base image
FROM golang:1.11-alpine

#take care of cache
RUN apk --no-cache add ca-certificates

RUN mkdir /ADDRESS /go/ADDRESS

COPY /getTXT.go /go/ADDRESS/getTXT.go

RUN cd /go/ADDRESS && go build -o getTXT

ENTRYPOINT ["/go/ADDRESS/getTXT"]