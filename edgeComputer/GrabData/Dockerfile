#This will act as container1 grabbing data with bash scripts
FROM golang:1.11-alpine AS builder
#FROM alpine:3.10
#FROM alpine
RUN apk --no-cache add ca-certificates
#RUN apk add --no-cache bash
RUN apk --no-cache add wget

#create MEGAGRONROCKS directory
RUN mkdir /MEGATRONROCKS

#create run wget command directory
RUN mkdir /MEGATRONROCKS/runWGET

#copy wow.sh script to MEGATRONROCKS
COPY /wow.sh /go/MEGATRONROCKS/runWGET/wow.sh

#install bash
RUN apk --no-cache add bash

ENTRYPOINT ["bash","/go/MEGATRONROCKS/runWGET/wow.sh"]