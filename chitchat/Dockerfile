FROM golang:1.18.3

RUN apt-get update && apt-get install -y git

WORKDIR /go/src
COPY ./src /go/src

CMD ["go", "run", "server.go"]