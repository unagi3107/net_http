FROM golang:1.18.3-alpine

RUN apt-get update && apt-get install -y git

WORKDIR /go/src
RUN mkdir /go/cert
COPY ./src /go/src
COPY ./cert /go/cert

CMD ["go", "run", "main.go"]